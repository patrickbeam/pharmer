package azure

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"github.com/Azure/azure-sdk-for-go/arm/network"
	"github.com/Azure/azure-sdk-for-go/arm/resources/resources"
	armstorage "github.com/Azure/azure-sdk-for-go/arm/storage"
	azstore "github.com/Azure/azure-sdk-for-go/storage"
	proto "github.com/appscode/api/kubernetes/v1beta1"
	"github.com/appscode/errors"
	"github.com/appscode/go/types"
	"github.com/appscode/pharmer/api"
	"github.com/appscode/pharmer/cloud"
)

func (cm *clusterManager) create(req *proto.ClusterCreateRequest) error {
	err := cm.initContext(req)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.ins, err = cloud.NewInstances(cm.ctx, cm.cluster)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.conn, err = NewConnector(cm.ctx, cm.cluster)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.ctx.Store().Clusters().SaveCluster(cm.cluster)

	defer func(releaseReservedIp bool) {
		if cm.cluster.Status.Phase == api.ClusterPhasePending {
			cm.cluster.Status.Phase = api.ClusterPhaseFailing
		}
		cm.ctx.Store().Clusters().SaveCluster(cm.cluster)
		cm.ctx.Store().Instances().SaveInstances(cm.ins.Instances)
		cm.ctx.Logger().Infof("Cluster %v is %v", cm.cluster.Name, cm.cluster.Status.Phase)
		if cm.cluster.Status.Phase != api.ClusterPhaseReady {
			cm.ctx.Logger().Infof("Cluster %v is deleting", cm.cluster.Name)
			cm.delete(&proto.ClusterDeleteRequest{
				Name:              cm.cluster.Name,
				ReleaseReservedIp: releaseReservedIp,
			})
		}
	}(cm.cluster.Spec.MasterReservedIP == "auto")

	// Common stuff
	_, err = cm.ensureResourceGroup()
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.ctx.Logger().Infof("Resource group %v in zone %v created", cm.namer.ResourceGroupName(), cm.cluster.Spec.Zone)
	as, err := cm.ensureAvailablitySet()
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.ctx.Logger().Infof("Availablity set %v created", cm.namer.AvailablitySetName())
	sa, err := cm.createStorageAccount()
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	vn, err := cm.ensureVirtualNetwork()
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	sg, err := cm.createNetworkSecurityGroup()
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	sn, err := cm.createSubnetID(&vn, &sg)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	// -------------------------------------------------------------------ASSETS
	im := &instanceManager{cluster: cm.cluster, conn: cm.conn, namer: cm.namer}

	masterPIP, err := im.createPublicIP(cm.namer.PublicIPName(cm.namer.MasterName()), network.Static)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.cluster.Spec.MasterReservedIP = types.String(masterPIP.IPAddress)
	cm.cluster.DetectApiServerURL()

	// @dipta
	if cm.cluster.Spec.MasterExternalIP == "" {
		cm.cluster.Spec.MasterExternalIP = cm.cluster.Spec.MasterReservedIP
	}

	// IP >>>>>>>>>>>>>>>>
	// TODO(tamal): if cluster.Spec.ctx.MasterReservedIP == "auto"
	//	name := cluster.Spec.ctx.KubernetesMasterName + "-pip"
	//	// cluster.Spec.ctx.MasterExternalIP = *ip.IPAddress
	//	cluster.Spec.ctx.MasterReservedIP = *ip.IPAddress
	//	// cluster.Spec.ctx.ApiServerUrl = "https://" + *ip.IPAddress

	err = cloud.GenClusterCerts(cm.ctx, cm.cluster)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	// needed for master start-up config
	if err = cm.ctx.Store().Clusters().SaveCluster(cm.cluster); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.UploadStartupConfig()

	// Master Stuff
	masterNIC, err := im.createNetworkInterface(cm.namer.NetworkInterfaceName(cm.cluster.Spec.KubernetesMasterName), sg, sn, network.Static, cm.cluster.Spec.MasterInternalIP, masterPIP)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	err = cm.createNetworkSecurityRule(&sg)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	masterScript := im.RenderStartupScript(cm.cluster.Spec.MasterSKU, api.RoleKubernetesMaster)
	masterVM, err := im.createVirtualMachine(masterNIC, as, sa, cm.namer.MasterName(), masterScript, cm.cluster.Spec.MasterSKU)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	ki, err := im.newKubeInstance(masterVM, masterNIC, masterPIP)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	ki.Spec.Role = api.RoleKubernetesMaster

	fmt.Println(cm.cluster.Spec.MasterExternalIP, "------------------------------->")
	cm.ins.Instances = append(cm.ins.Instances, ki)

	err = cloud.EnsureARecord(cm.ctx, cm.cluster, ki) // works for reserved or non-reserved mode
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	fmt.Println(err, "<------------------------------->")

	for _, ng := range req.NodeGroups {
		cm.ctx.Logger().Infof("Creating %v node with sku %v", ng.Count, ng.Sku)
		igm := &InstanceGroupManager{
			cm: cm,
			instance: cloud.Instance{
				Type: cloud.InstanceType{
					ContextVersion: cm.cluster.Spec.ResourceVersion,
					Sku:            ng.Sku,

					Master:       false,
					SpotInstance: false,
				},
				Stats: cloud.GroupStats{
					Count: ng.Count,
				},
			},
			im: im,
		}
		err = igm.AdjustInstanceGroup()

	}

	cm.ctx.Logger().Info("Waiting for cluster initialization")

	// Wait for master A record to propagate
	if err := cloud.EnsureDnsIPLookup(cm.ctx, cm.cluster); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	// wait for nodes to start
	if err := cloud.ProbeKubeAPI(cm.ctx, cm.cluster); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	// check all components are ok
	if err = cloud.CheckComponentStatuses(cm.ctx, cm.cluster); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	// Make sure nodes are connected to master and are ready
	if err = cloud.WaitForReadyNodes(cm.ctx, cm.cluster); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	cm.cluster.Status.Phase = api.ClusterPhaseReady
	return nil
}

// IP >>>>>>>>>>>>>>>>
// TODO(tamal): if cluster.Spec.ctx.MasterReservedIP == "auto"
//	name := cluster.Spec.ctx.KubernetesMasterName + "-pip"
//	// cluster.Spec.ctx.MasterExternalIP = *ip.IPAddress
//	cluster.Spec.ctx.MasterReservedIP = *ip.IPAddress
//	// cluster.Spec.ctx.ApiServerUrl = "https://" + *ip.IPAddress

func (cm *clusterManager) ensureResourceGroup() (resources.Group, error) {
	req := resources.Group{
		Name:     types.StringP(cm.namer.ResourceGroupName()),
		Location: types.StringP(cm.cluster.Spec.Zone),
		Tags: &map[string]*string{
			"KubernetesCluster": types.StringP(cm.cluster.Name),
		},
	}
	return cm.conn.groupsClient.CreateOrUpdate(cm.namer.ResourceGroupName(), req)
}

func (cm *clusterManager) ensureAvailablitySet() (compute.AvailabilitySet, error) {
	name := cm.namer.AvailablitySetName()
	req := compute.AvailabilitySet{
		Name:     types.StringP(name),
		Location: types.StringP(cm.cluster.Spec.Zone),
		Tags: &map[string]*string{
			"KubernetesCluster": types.StringP(cm.cluster.Name),
		},
	}
	return cm.conn.availabilitySetsClient.CreateOrUpdate(cm.namer.ResourceGroupName(), name, req)
}

func (cm *clusterManager) ensureVirtualNetwork() (network.VirtualNetwork, error) {
	name := cm.namer.VirtualNetworkName()
	req := network.VirtualNetwork{
		Name:     types.StringP(name),
		Location: types.StringP(cm.cluster.Spec.Zone),
		VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: &[]string{cm.cluster.Spec.NonMasqueradeCidr},
			},
		},
		Tags: &map[string]*string{
			"KubernetesCluster": types.StringP(cm.cluster.Name),
		},
	}

	_, errchan := cm.conn.virtualNetworksClient.CreateOrUpdate(cm.namer.ResourceGroupName(), name, req, nil)
	err := <-errchan
	if err != nil {
		return network.VirtualNetwork{}, err
	}
	cm.ctx.Logger().Infof("Virtual network %v created", name)
	return cm.conn.virtualNetworksClient.Get(cm.namer.ResourceGroupName(), name, "")
}

func (cm *clusterManager) getVirtualNetwork() (network.VirtualNetwork, error) {
	return cm.conn.virtualNetworksClient.Get(cm.namer.ResourceGroupName(), cm.namer.VirtualNetworkName(), "")
}

func (cm *clusterManager) createNetworkSecurityGroup() (network.SecurityGroup, error) {
	securityGroupName := cm.namer.NetworkSecurityGroupName()
	securityGroup := network.SecurityGroup{
		Name:     types.StringP(securityGroupName),
		Location: types.StringP(cm.cluster.Spec.Zone),
		Tags: &map[string]*string{
			"KubernetesCluster": types.StringP(cm.cluster.Name),
		},
	}
	_, errchan := cm.conn.securityGroupsClient.CreateOrUpdate(cm.namer.ResourceGroupName(), securityGroupName, securityGroup, nil)
	err := <-errchan
	if err != nil {
		return network.SecurityGroup{}, err
	}
	cm.ctx.Logger().Infof("Network security group %v created", securityGroupName)
	return cm.conn.securityGroupsClient.Get(cm.namer.ResourceGroupName(), securityGroupName, "")
}

func (cm *clusterManager) getNetworkSecurityGroup() (network.SecurityGroup, error) {
	securityGroupName := cm.namer.NetworkSecurityGroupName()
	return cm.conn.securityGroupsClient.Get(cm.namer.ResourceGroupName(), securityGroupName, "")
}

func (cm *clusterManager) createSubnetID(vn *network.VirtualNetwork, sg *network.SecurityGroup) (network.Subnet, error) {
	name := cm.namer.SubnetName()
	routeTable, err := cm.createRouteTable()
	if err != nil {
		return network.Subnet{}, err
	}
	req := network.Subnet{
		Name: types.StringP(name),
		SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
			NetworkSecurityGroup: &network.SecurityGroup{
				ID: sg.ID,
			},
			AddressPrefix: types.StringP(cm.cluster.Spec.SubnetCidr),
			RouteTable: &network.RouteTable{
				ID: routeTable.ID,
			},
		},
	}

	_, errchan := cm.conn.subnetsClient.CreateOrUpdate(cm.namer.ResourceGroupName(), *vn.Name, name, req, nil)
	err = <-errchan
	if err != nil {
		return network.Subnet{}, err
	}
	cm.ctx.Logger().Infof("Subnet name %v created", name)
	return cm.conn.subnetsClient.Get(cm.namer.ResourceGroupName(), *vn.Name, name, "")
}

func (cm *clusterManager) getSubnetID(vn *network.VirtualNetwork) (network.Subnet, error) {
	return cm.conn.subnetsClient.Get(cm.namer.ResourceGroupName(), *vn.Name, cm.namer.SubnetName(), "")
}

func (cm *clusterManager) createRouteTable() (network.RouteTable, error) {
	name := cm.namer.RouteTableName()
	req := network.RouteTable{
		Name:     types.StringP(name),
		Location: types.StringP(cm.cluster.Spec.Zone),
		Tags: &map[string]*string{
			"KubernetesCluster": types.StringP(cm.cluster.Name),
		},
	}
	_, errchan := cm.conn.routeTablesClient.CreateOrUpdate(cm.namer.ResourceGroupName(), name, req, nil)
	err := <-errchan
	if err != nil {
		return network.RouteTable{}, err
	}
	cm.ctx.Logger().Infof("Route table %v created", name)
	return cm.conn.routeTablesClient.Get(cm.namer.ResourceGroupName(), name, "")
}

func (cm *clusterManager) createNetworkSecurityRule(sg *network.SecurityGroup) error {
	sshRuleName := cm.namer.NetworkSecurityRule("ssh")
	sshRule := network.SecurityRule{
		Name: types.StringP(sshRuleName),
		SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
			Access: network.SecurityRuleAccessAllow,
			DestinationAddressPrefix: types.StringP("*"),
			DestinationPortRange:     types.StringP("22"),
			Direction:                network.SecurityRuleDirectionInbound,
			Priority:                 types.Int32P(100),
			Protocol:                 network.SecurityRuleProtocolTCP,
			SourceAddressPrefix:      types.StringP("*"),
			SourcePortRange:          types.StringP("*"),
		},
	}
	_, errchan := cm.conn.securityRulesClient.CreateOrUpdate(cm.namer.ResourceGroupName(), *sg.Name, sshRuleName, sshRule, nil)
	err := <-errchan
	if err != nil {
		return err
	}
	cm.ctx.Logger().Infof("Network security rule %v created", sshRuleName)
	sslRuleName := cm.namer.NetworkSecurityRule("ssl")
	sslRule := network.SecurityRule{
		Name: types.StringP(sshRuleName),
		SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
			Access: network.SecurityRuleAccessAllow,
			DestinationAddressPrefix: types.StringP("*"),
			DestinationPortRange:     types.StringP("443"),
			Direction:                network.SecurityRuleDirectionInbound,
			Priority:                 types.Int32P(110),
			Protocol:                 network.SecurityRuleProtocolTCP,
			SourceAddressPrefix:      types.StringP("*"),
			SourcePortRange:          types.StringP("*"),
		},
	}
	_, errchan = cm.conn.securityRulesClient.CreateOrUpdate(cm.namer.ResourceGroupName(), *sg.Name, sslRuleName, sslRule, nil)
	err = <-errchan
	if err != nil {
		return err
	}
	cm.ctx.Logger().Infof("Network security rule %v created", sslRuleName)

	mastersslRuleName := cm.namer.NetworkSecurityRule("masterssl")
	mastersslRule := network.SecurityRule{
		Name: types.StringP(mastersslRuleName),
		SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
			Access: network.SecurityRuleAccessAllow,
			DestinationAddressPrefix: types.StringP("*"),
			DestinationPortRange:     types.StringP("6443"),
			Direction:                network.SecurityRuleDirectionInbound,
			Priority:                 types.Int32P(120),
			Protocol:                 network.SecurityRuleProtocolTCP,
			SourceAddressPrefix:      types.StringP("*"),
			SourcePortRange:          types.StringP("*"),
		},
	}
	_, errchan = cm.conn.securityRulesClient.CreateOrUpdate(cm.namer.ResourceGroupName(), *sg.Name, mastersslRuleName, mastersslRule, nil)
	err = <-errchan
	if err != nil {
		return err
	}
	cm.ctx.Logger().Infof("Network security rule %v created", mastersslRuleName)

	return err
}

func (cm *clusterManager) createStorageAccount() (armstorage.Account, error) {
	storageName := cm.cluster.Spec.AzureCloudConfig.StorageAccountName
	req := armstorage.AccountCreateParameters{
		Location: types.StringP(cm.cluster.Spec.Zone),
		Sku: &armstorage.Sku{
			Name: armstorage.StandardLRS,
		},
		Tags: &map[string]*string{
			"KubernetesCluster": types.StringP(cm.cluster.Name),
		},
	}
	_, errchan := cm.conn.storageClient.Create(cm.namer.ResourceGroupName(), storageName, req, nil)
	err := <-errchan
	if err != nil {
		return armstorage.Account{}, err
	}
	cm.ctx.Logger().Infof("Storage account %v created", storageName)
	keys, err := cm.conn.storageClient.ListKeys(cm.namer.ResourceGroupName(), storageName)
	if err != nil {
		return armstorage.Account{}, err
	}
	storageClient, err := azstore.NewBasicClient(storageName, *(*(keys.Keys))[0].Value)
	if err != nil {
		return armstorage.Account{}, err
	}

	bs := storageClient.GetBlobService()
	_, err = bs.GetContainerReference(cm.namer.StorageContainerName()).CreateIfNotExists(&azstore.CreateContainerOptions{Access: azstore.ContainerAccessTypeContainer})
	if err != nil {
		return armstorage.Account{}, err
	}
	return cm.conn.storageClient.GetProperties(cm.namer.ResourceGroupName(), storageName)
}

/*
for i := int64(0); i < ng.Count; i++ {
			nodeName := cm.namer.GenNodeName(ng.Sku)

			nodePIP, err := im.createPublicIP(cm.namer.PublicIPName(nodeName), network.Dynamic)
			if err != nil {
				cm.ctx.StatusCause = err.Error()
				return errors.FromErr(err).WithContext(cm.ctx).Err()
			}

			nodeNIC, err := im.createNetworkInterface(cm.namer.NetworkInterfaceName(nodeName), sn, network.Dynamic, "", nodePIP)
			if err != nil {
				cm.ctx.StatusCause = err.Error()
				return errors.FromErr(err).WithContext(cm.ctx).Err()
			}

			nodeScript := im.RenderStartupScript(cm.ctx.NewScriptOptions(), ng.Sku, api.RoleKubernetesPool)
			nodeVM, err := im.createVirtualMachine(nodeNIC, as, sa, nodeName, nodeScript, ng.Sku)
			if err != nil {
				cm.ctx.StatusCause = err.Error()
				return errors.FromErr(err).WithContext(cm.ctx).Err()
			}

			nodePIP, err = im.getPublicIP(cm.namer.PublicIPName(nodeName))
			if err != nil {
				cm.ctx.StatusCause = err.Error()
				return errors.FromErr(err).WithContext(cm.ctx).Err()
			}

			ki, err := im.newKubeInstance(nodeVM, nodeNIC, nodePIP)
			if err != nil {
				cm.ctx.StatusCause = err.Error()
				return errors.FromErr(err).WithContext(cm.ctx).Err()
			}
			ki.Role = api.RoleKubernetesPool
			cm.ins.Instances = append(cm.ins.Instances, ki)
			// cm.ins.Instances = append(cm.ins.Instances, ki)
		}
*/
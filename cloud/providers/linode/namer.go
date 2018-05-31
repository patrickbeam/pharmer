package linode

import (
	"github.com/appscode/go/crypto/rand"
	api "github.com/pharmer/pharmer/apis/v1"
)

type namer struct {
	cluster *api.Cluster
}

func (n namer) MasterName() string {
	return n.cluster.Name + "-master"
}

// Deprecated
func (n namer) GenNodeName(ng string) string {
	return rand.WithUniqSuffix(ng)
}

func (n namer) GenSSHKeyExternalID() string {
	return n.cluster.Name + "-" + rand.Characters(6)
}

func (n namer) StartupScriptName(ng, role string) string {
	return ng + "-" + role
}

func (n namer) LoadBalancerName() string {
	return n.cluster.Name + "-lb"
}

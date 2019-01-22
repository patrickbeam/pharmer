package digitalocean

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	clusterapi "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"sigs.k8s.io/cluster-api/pkg/controller/cluster"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, func(ctx context.Context, m manager.Manager) error {
		actuator := NewClusterActuator(m, ClusterActuatorParams{
			Ctx:           ctx,
			EventRecorder: m.GetRecorder(Recorder),
			Scheme:        m.GetScheme(),
		})
		return cluster.AddWithActuator(m, actuator)
	})

}

type ClusterActuator struct {
	ctx           context.Context
	client        client.Client
	eventRecorder record.EventRecorder
	scheme        *runtime.Scheme
}

type ClusterActuatorParams struct {
	Ctx           context.Context
	EventRecorder record.EventRecorder
	Scheme        *runtime.Scheme
}

func NewClusterActuator(m manager.Manager, params ClusterActuatorParams) *ClusterActuator {
	return &ClusterActuator{
		ctx:           params.Ctx,
		client:        m.GetClient(),
		eventRecorder: params.EventRecorder,
		scheme:        params.Scheme,
	}
}

/*func (cm *ClusterManager) InitializeActuator(client client.ClusterV1alpha1Interface, rec record.EventRecorder, scheme *runtime.Scheme) error {
	cm.actuator = &ClusterActuator{
		client:        client,
		eventRecorder: rec,
		scheme:        scheme,
	}
	fmt.Println("ClusterActuator initialized")

	return nil
}*/

func (cm *ClusterActuator) Reconcile(cluster *clusterapi.Cluster) error {
	fmt.Println("Reconciling cluster %v", cluster.Name)

	return nil
}

func (cm *ClusterActuator) Delete(cluster *clusterapi.Cluster) error {
	fmt.Println("Delete cluster %v", cluster.Name)
	return nil
}
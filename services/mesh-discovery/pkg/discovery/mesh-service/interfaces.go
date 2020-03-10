package mesh_service

import (
	"github.com/solo-io/mesh-projects/pkg/api/discovery.zephyr.solo.io/v1alpha1/controller"
	corev1_controllers "github.com/solo-io/mesh-projects/services/common/cluster/core/v1/controller"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go -package service_discovery_mocks

type MeshServiceFinder interface {
	StartDiscovery(serviceController corev1_controllers.ServiceController, meshWorkloadController controller.MeshWorkloadController) error
}
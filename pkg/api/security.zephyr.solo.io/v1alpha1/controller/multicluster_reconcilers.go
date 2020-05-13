// Code generated by skv2. DO NOT EDIT.

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	security_zephyr_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the VirtualMeshCertificateSigningRequest Resource across clusters.
// implemented by the user
type MulticlusterVirtualMeshCertificateSigningRequestReconciler interface {
	ReconcileVirtualMeshCertificateSigningRequest(clusterName string, obj *security_zephyr_solo_io_v1alpha1.VirtualMeshCertificateSigningRequest) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualMeshCertificateSigningRequest Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler interface {
	ReconcileVirtualMeshCertificateSigningRequestDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterVirtualMeshCertificateSigningRequestReconcilerFuncs struct {
	OnReconcileVirtualMeshCertificateSigningRequest         func(clusterName string, obj *security_zephyr_solo_io_v1alpha1.VirtualMeshCertificateSigningRequest) (reconcile.Result, error)
	OnReconcileVirtualMeshCertificateSigningRequestDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterVirtualMeshCertificateSigningRequestReconcilerFuncs) ReconcileVirtualMeshCertificateSigningRequest(clusterName string, obj *security_zephyr_solo_io_v1alpha1.VirtualMeshCertificateSigningRequest) (reconcile.Result, error) {
	if f.OnReconcileVirtualMeshCertificateSigningRequest == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualMeshCertificateSigningRequest(clusterName, obj)
}

func (f *MulticlusterVirtualMeshCertificateSigningRequestReconcilerFuncs) ReconcileVirtualMeshCertificateSigningRequestDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileVirtualMeshCertificateSigningRequestDeletion == nil {
		return
	}
	f.OnReconcileVirtualMeshCertificateSigningRequestDeletion(clusterName, req)
}

type MulticlusterVirtualMeshCertificateSigningRequestReconcileLoop interface {
	// AddMulticlusterVirtualMeshCertificateSigningRequestReconciler adds a MulticlusterVirtualMeshCertificateSigningRequestReconciler to the MulticlusterVirtualMeshCertificateSigningRequestReconcileLoop.
	AddMulticlusterVirtualMeshCertificateSigningRequestReconciler(ctx context.Context, rec MulticlusterVirtualMeshCertificateSigningRequestReconciler, predicates ...predicate.Predicate)
}

type multiclusterVirtualMeshCertificateSigningRequestReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterVirtualMeshCertificateSigningRequestReconcileLoop) AddMulticlusterVirtualMeshCertificateSigningRequestReconciler(ctx context.Context, rec MulticlusterVirtualMeshCertificateSigningRequestReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericVirtualMeshCertificateSigningRequestMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterVirtualMeshCertificateSigningRequestReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterVirtualMeshCertificateSigningRequestReconcileLoop {
	return &multiclusterVirtualMeshCertificateSigningRequestReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_zephyr_solo_io_v1alpha1.VirtualMeshCertificateSigningRequest{})}
}

type genericVirtualMeshCertificateSigningRequestMulticlusterReconciler struct {
	reconciler MulticlusterVirtualMeshCertificateSigningRequestReconciler
}

func (g genericVirtualMeshCertificateSigningRequestMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterVirtualMeshCertificateSigningRequestDeletionReconciler); ok {
		deletionReconciler.ReconcileVirtualMeshCertificateSigningRequestDeletion(cluster, req)
	}
}

func (g genericVirtualMeshCertificateSigningRequestMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_zephyr_solo_io_v1alpha1.VirtualMeshCertificateSigningRequest)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualMeshCertificateSigningRequest handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualMeshCertificateSigningRequest(cluster, obj)
}

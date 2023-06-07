//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KubeStellar Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	edgev1alpha1 "github.com/kcp-dev/edge-mc/pkg/apis/edge/v1alpha1"
	scopedclientset "github.com/kcp-dev/edge-mc/pkg/client/clientset/versioned"
	clientset "github.com/kcp-dev/edge-mc/pkg/client/clientset/versioned/cluster"
	"github.com/kcp-dev/edge-mc/pkg/client/informers/externalversions/internalinterfaces"
	edgev1alpha1listers "github.com/kcp-dev/edge-mc/pkg/client/listers/edge/v1alpha1"
)

// LogicalClusterClusterInformer provides access to a shared informer and lister for
// LogicalClusters.
type LogicalClusterClusterInformer interface {
	Cluster(logicalcluster.Name) LogicalClusterInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() edgev1alpha1listers.LogicalClusterClusterLister
}

type logicalClusterClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewLogicalClusterClusterInformer constructs a new informer for LogicalCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLogicalClusterClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredLogicalClusterClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredLogicalClusterClusterInformer constructs a new informer for LogicalCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLogicalClusterClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdgeV1alpha1().LogicalClusters().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdgeV1alpha1().LogicalClusters().Watch(context.TODO(), options)
			},
		},
		&edgev1alpha1.LogicalCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *logicalClusterClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredLogicalClusterClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *logicalClusterClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&edgev1alpha1.LogicalCluster{}, f.defaultInformer)
}

func (f *logicalClusterClusterInformer) Lister() edgev1alpha1listers.LogicalClusterClusterLister {
	return edgev1alpha1listers.NewLogicalClusterClusterLister(f.Informer().GetIndexer())
}

// LogicalClusterInformer provides access to a shared informer and lister for
// LogicalClusters.
type LogicalClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() edgev1alpha1listers.LogicalClusterLister
}

func (f *logicalClusterClusterInformer) Cluster(clusterName logicalcluster.Name) LogicalClusterInformer {
	return &logicalClusterInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type logicalClusterInformer struct {
	informer cache.SharedIndexInformer
	lister   edgev1alpha1listers.LogicalClusterLister
}

func (f *logicalClusterInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *logicalClusterInformer) Lister() edgev1alpha1listers.LogicalClusterLister {
	return f.lister
}

type logicalClusterScopedInformer struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func (f *logicalClusterScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&edgev1alpha1.LogicalCluster{}, f.defaultInformer)
}

func (f *logicalClusterScopedInformer) Lister() edgev1alpha1listers.LogicalClusterLister {
	return edgev1alpha1listers.NewLogicalClusterLister(f.Informer().GetIndexer())
}

// NewLogicalClusterInformer constructs a new informer for LogicalCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLogicalClusterInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLogicalClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredLogicalClusterInformer constructs a new informer for LogicalCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLogicalClusterInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdgeV1alpha1().LogicalClusters().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdgeV1alpha1().LogicalClusters().Watch(context.TODO(), options)
			},
		},
		&edgev1alpha1.LogicalCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *logicalClusterScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLogicalClusterInformer(client, resyncPeriod, cache.Indexers{}, f.tweakListOptions)
}
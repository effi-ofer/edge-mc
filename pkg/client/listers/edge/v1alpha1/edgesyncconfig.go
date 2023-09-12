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


//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"	
	"github.com/kcp-dev/logicalcluster/v3"
	
	"k8s.io/client-go/tools/cache"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/api/errors"

	edgev1alpha1 "github.com/kubestellar/kubestellar/pkg/apis/edge/v1alpha1"
	)

// EdgeSyncConfigClusterLister can list EdgeSyncConfigs across all workspaces, or scope down to a EdgeSyncConfigLister for one workspace.
// All objects returned here must be treated as read-only.
type EdgeSyncConfigClusterLister interface {
	// List lists all EdgeSyncConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*edgev1alpha1.EdgeSyncConfig, err error)
	// Cluster returns a lister that can list and get EdgeSyncConfigs in one workspace.
Cluster(clusterName logicalcluster.Name) EdgeSyncConfigLister
EdgeSyncConfigClusterListerExpansion
}

type edgeSyncConfigClusterLister struct {
	indexer cache.Indexer
}

// NewEdgeSyncConfigClusterLister returns a new EdgeSyncConfigClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewEdgeSyncConfigClusterLister(indexer cache.Indexer) *edgeSyncConfigClusterLister {
	return &edgeSyncConfigClusterLister{indexer: indexer}
}

// List lists all EdgeSyncConfigs in the indexer across all workspaces.
func (s *edgeSyncConfigClusterLister) List(selector labels.Selector) (ret []*edgev1alpha1.EdgeSyncConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*edgev1alpha1.EdgeSyncConfig))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get EdgeSyncConfigs.
func (s *edgeSyncConfigClusterLister) Cluster(clusterName logicalcluster.Name) EdgeSyncConfigLister {
return &edgeSyncConfigLister{indexer: s.indexer, clusterName: clusterName}
}

// EdgeSyncConfigLister can list all EdgeSyncConfigs, or get one in particular.
// All objects returned here must be treated as read-only.
type EdgeSyncConfigLister interface {
	// List lists all EdgeSyncConfigs in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*edgev1alpha1.EdgeSyncConfig, err error)
// Get retrieves the EdgeSyncConfig from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*edgev1alpha1.EdgeSyncConfig, error)
EdgeSyncConfigListerExpansion
}
// edgeSyncConfigLister can list all EdgeSyncConfigs inside a workspace.
type edgeSyncConfigLister struct {
	indexer cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all EdgeSyncConfigs in the indexer for a workspace.
func (s *edgeSyncConfigLister) List(selector labels.Selector) (ret []*edgev1alpha1.EdgeSyncConfig, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*edgev1alpha1.EdgeSyncConfig))
	})
	return ret, err
}

// Get retrieves the EdgeSyncConfig from the indexer for a given workspace and name.
func (s *edgeSyncConfigLister) Get(name string) (*edgev1alpha1.EdgeSyncConfig, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(edgev1alpha1.Resource("EdgeSyncConfig"), name)
	}
	return obj.(*edgev1alpha1.EdgeSyncConfig), nil
}
// NewEdgeSyncConfigLister returns a new EdgeSyncConfigLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewEdgeSyncConfigLister(indexer cache.Indexer) *edgeSyncConfigScopedLister {
	return &edgeSyncConfigScopedLister{indexer: indexer}
}

// edgeSyncConfigScopedLister can list all EdgeSyncConfigs inside a workspace.
type edgeSyncConfigScopedLister struct {
	indexer cache.Indexer
}

// List lists all EdgeSyncConfigs in the indexer for a workspace.
func (s *edgeSyncConfigScopedLister) List(selector labels.Selector) (ret []*edgev1alpha1.EdgeSyncConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*edgev1alpha1.EdgeSyncConfig))
	})
	return ret, err
}

// Get retrieves the EdgeSyncConfig from the indexer for a given workspace and name.
func (s *edgeSyncConfigScopedLister) Get(name string) (*edgev1alpha1.EdgeSyncConfig, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(edgev1alpha1.Resource("EdgeSyncConfig"), name)
	}
	return obj.(*edgev1alpha1.EdgeSyncConfig), nil
}

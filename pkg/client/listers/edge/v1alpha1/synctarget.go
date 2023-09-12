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

// SyncTargetClusterLister can list SyncTargets across all workspaces, or scope down to a SyncTargetLister for one workspace.
// All objects returned here must be treated as read-only.
type SyncTargetClusterLister interface {
	// List lists all SyncTargets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*edgev1alpha1.SyncTarget, err error)
	// Cluster returns a lister that can list and get SyncTargets in one workspace.
Cluster(clusterName logicalcluster.Name) SyncTargetLister
SyncTargetClusterListerExpansion
}

type syncTargetClusterLister struct {
	indexer cache.Indexer
}

// NewSyncTargetClusterLister returns a new SyncTargetClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewSyncTargetClusterLister(indexer cache.Indexer) *syncTargetClusterLister {
	return &syncTargetClusterLister{indexer: indexer}
}

// List lists all SyncTargets in the indexer across all workspaces.
func (s *syncTargetClusterLister) List(selector labels.Selector) (ret []*edgev1alpha1.SyncTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*edgev1alpha1.SyncTarget))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get SyncTargets.
func (s *syncTargetClusterLister) Cluster(clusterName logicalcluster.Name) SyncTargetLister {
return &syncTargetLister{indexer: s.indexer, clusterName: clusterName}
}

// SyncTargetLister can list all SyncTargets, or get one in particular.
// All objects returned here must be treated as read-only.
type SyncTargetLister interface {
	// List lists all SyncTargets in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*edgev1alpha1.SyncTarget, err error)
// Get retrieves the SyncTarget from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*edgev1alpha1.SyncTarget, error)
SyncTargetListerExpansion
}
// syncTargetLister can list all SyncTargets inside a workspace.
type syncTargetLister struct {
	indexer cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all SyncTargets in the indexer for a workspace.
func (s *syncTargetLister) List(selector labels.Selector) (ret []*edgev1alpha1.SyncTarget, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*edgev1alpha1.SyncTarget))
	})
	return ret, err
}

// Get retrieves the SyncTarget from the indexer for a given workspace and name.
func (s *syncTargetLister) Get(name string) (*edgev1alpha1.SyncTarget, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(edgev1alpha1.Resource("SyncTarget"), name)
	}
	return obj.(*edgev1alpha1.SyncTarget), nil
}
// NewSyncTargetLister returns a new SyncTargetLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewSyncTargetLister(indexer cache.Indexer) *syncTargetScopedLister {
	return &syncTargetScopedLister{indexer: indexer}
}

// syncTargetScopedLister can list all SyncTargets inside a workspace.
type syncTargetScopedLister struct {
	indexer cache.Indexer
}

// List lists all SyncTargets in the indexer for a workspace.
func (s *syncTargetScopedLister) List(selector labels.Selector) (ret []*edgev1alpha1.SyncTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*edgev1alpha1.SyncTarget))
	})
	return ret, err
}

// Get retrieves the SyncTarget from the indexer for a given workspace and name.
func (s *syncTargetScopedLister) Get(name string) (*edgev1alpha1.SyncTarget, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(edgev1alpha1.Resource("SyncTarget"), name)
	}
	return obj.(*edgev1alpha1.SyncTarget), nil
}

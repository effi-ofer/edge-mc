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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	edgev1alpha1 "github.com/kcp-dev/edge-mc/pkg/apis/edge/v1alpha1"
)

// ClusterProviderDescClusterLister can list ClusterProviderDescs across all workspaces, or scope down to a ClusterProviderDescLister for one workspace.
// All objects returned here must be treated as read-only.
type ClusterProviderDescClusterLister interface {
	// List lists all ClusterProviderDescs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*edgev1alpha1.ClusterProviderDesc, err error)
	// Cluster returns a lister that can list and get ClusterProviderDescs in one workspace.
	Cluster(clusterName logicalcluster.Name) ClusterProviderDescLister
	ClusterProviderDescClusterListerExpansion
}

type clusterProviderDescClusterLister struct {
	indexer cache.Indexer
}

// NewClusterProviderDescClusterLister returns a new ClusterProviderDescClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewClusterProviderDescClusterLister(indexer cache.Indexer) *clusterProviderDescClusterLister {
	return &clusterProviderDescClusterLister{indexer: indexer}
}

// List lists all ClusterProviderDescs in the indexer across all workspaces.
func (s *clusterProviderDescClusterLister) List(selector labels.Selector) (ret []*edgev1alpha1.ClusterProviderDesc, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*edgev1alpha1.ClusterProviderDesc))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ClusterProviderDescs.
func (s *clusterProviderDescClusterLister) Cluster(clusterName logicalcluster.Name) ClusterProviderDescLister {
	return &clusterProviderDescLister{indexer: s.indexer, clusterName: clusterName}
}

// ClusterProviderDescLister can list all ClusterProviderDescs, or get one in particular.
// All objects returned here must be treated as read-only.
type ClusterProviderDescLister interface {
	// List lists all ClusterProviderDescs in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*edgev1alpha1.ClusterProviderDesc, err error)
	// Get retrieves the ClusterProviderDesc from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*edgev1alpha1.ClusterProviderDesc, error)
	ClusterProviderDescListerExpansion
}

// clusterProviderDescLister can list all ClusterProviderDescs inside a workspace.
type clusterProviderDescLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all ClusterProviderDescs in the indexer for a workspace.
func (s *clusterProviderDescLister) List(selector labels.Selector) (ret []*edgev1alpha1.ClusterProviderDesc, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*edgev1alpha1.ClusterProviderDesc))
	})
	return ret, err
}

// Get retrieves the ClusterProviderDesc from the indexer for a given workspace and name.
func (s *clusterProviderDescLister) Get(name string) (*edgev1alpha1.ClusterProviderDesc, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(edgev1alpha1.Resource("ClusterProviderDesc"), name)
	}
	return obj.(*edgev1alpha1.ClusterProviderDesc), nil
}

// NewClusterProviderDescLister returns a new ClusterProviderDescLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewClusterProviderDescLister(indexer cache.Indexer) *clusterProviderDescScopedLister {
	return &clusterProviderDescScopedLister{indexer: indexer}
}

// clusterProviderDescScopedLister can list all ClusterProviderDescs inside a workspace.
type clusterProviderDescScopedLister struct {
	indexer cache.Indexer
}

// List lists all ClusterProviderDescs in the indexer for a workspace.
func (s *clusterProviderDescScopedLister) List(selector labels.Selector) (ret []*edgev1alpha1.ClusterProviderDesc, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*edgev1alpha1.ClusterProviderDesc))
	})
	return ret, err
}

// Get retrieves the ClusterProviderDesc from the indexer for a given workspace and name.
func (s *clusterProviderDescScopedLister) Get(name string) (*edgev1alpha1.ClusterProviderDesc, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(edgev1alpha1.Resource("ClusterProviderDesc"), name)
	}
	return obj.(*edgev1alpha1.ClusterProviderDesc), nil
}

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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"github.com/kcp-dev/logicalcluster/v3"

	logicalclusterv1alpha1 "github.com/kcp-dev/edge-mc/pkg/apis/logicalcluster/v1alpha1"
	logicalclusterv1alpha1client "github.com/kcp-dev/edge-mc/pkg/client/clientset/versioned/typed/logicalcluster/v1alpha1"
)

var clusterProviderDescsResource = schema.GroupVersionResource{Group: "logicalcluster.kubestellar.io", Version: "v1alpha1", Resource: "clusterproviderdescs"}
var clusterProviderDescsKind = schema.GroupVersionKind{Group: "logicalcluster.kubestellar.io", Version: "v1alpha1", Kind: "ClusterProviderDesc"}

type clusterProviderDescsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterProviderDescsClusterClient) Cluster(clusterPath logicalcluster.Path) logicalclusterv1alpha1client.ClusterProviderDescInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &clusterProviderDescsClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ClusterProviderDescs that match those selectors across all clusters.
func (c *clusterProviderDescsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*logicalclusterv1alpha1.ClusterProviderDescList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterProviderDescsResource, clusterProviderDescsKind, logicalcluster.Wildcard, opts), &logicalclusterv1alpha1.ClusterProviderDescList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &logicalclusterv1alpha1.ClusterProviderDescList{ListMeta: obj.(*logicalclusterv1alpha1.ClusterProviderDescList).ListMeta}
	for _, item := range obj.(*logicalclusterv1alpha1.ClusterProviderDescList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ClusterProviderDescs across all clusters.
func (c *clusterProviderDescsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterProviderDescsResource, logicalcluster.Wildcard, opts))
}

type clusterProviderDescsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *clusterProviderDescsClient) Create(ctx context.Context, clusterProviderDesc *logicalclusterv1alpha1.ClusterProviderDesc, opts metav1.CreateOptions) (*logicalclusterv1alpha1.ClusterProviderDesc, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(clusterProviderDescsResource, c.ClusterPath, clusterProviderDesc), &logicalclusterv1alpha1.ClusterProviderDesc{})
	if obj == nil {
		return nil, err
	}
	return obj.(*logicalclusterv1alpha1.ClusterProviderDesc), err
}

func (c *clusterProviderDescsClient) Update(ctx context.Context, clusterProviderDesc *logicalclusterv1alpha1.ClusterProviderDesc, opts metav1.UpdateOptions) (*logicalclusterv1alpha1.ClusterProviderDesc, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(clusterProviderDescsResource, c.ClusterPath, clusterProviderDesc), &logicalclusterv1alpha1.ClusterProviderDesc{})
	if obj == nil {
		return nil, err
	}
	return obj.(*logicalclusterv1alpha1.ClusterProviderDesc), err
}

func (c *clusterProviderDescsClient) UpdateStatus(ctx context.Context, clusterProviderDesc *logicalclusterv1alpha1.ClusterProviderDesc, opts metav1.UpdateOptions) (*logicalclusterv1alpha1.ClusterProviderDesc, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(clusterProviderDescsResource, c.ClusterPath, "status", clusterProviderDesc), &logicalclusterv1alpha1.ClusterProviderDesc{})
	if obj == nil {
		return nil, err
	}
	return obj.(*logicalclusterv1alpha1.ClusterProviderDesc), err
}

func (c *clusterProviderDescsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(clusterProviderDescsResource, c.ClusterPath, name, opts), &logicalclusterv1alpha1.ClusterProviderDesc{})
	return err
}

func (c *clusterProviderDescsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(clusterProviderDescsResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &logicalclusterv1alpha1.ClusterProviderDescList{})
	return err
}

func (c *clusterProviderDescsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*logicalclusterv1alpha1.ClusterProviderDesc, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(clusterProviderDescsResource, c.ClusterPath, name), &logicalclusterv1alpha1.ClusterProviderDesc{})
	if obj == nil {
		return nil, err
	}
	return obj.(*logicalclusterv1alpha1.ClusterProviderDesc), err
}

// List takes label and field selectors, and returns the list of ClusterProviderDescs that match those selectors.
func (c *clusterProviderDescsClient) List(ctx context.Context, opts metav1.ListOptions) (*logicalclusterv1alpha1.ClusterProviderDescList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterProviderDescsResource, clusterProviderDescsKind, c.ClusterPath, opts), &logicalclusterv1alpha1.ClusterProviderDescList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &logicalclusterv1alpha1.ClusterProviderDescList{ListMeta: obj.(*logicalclusterv1alpha1.ClusterProviderDescList).ListMeta}
	for _, item := range obj.(*logicalclusterv1alpha1.ClusterProviderDescList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *clusterProviderDescsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterProviderDescsResource, c.ClusterPath, opts))
}

func (c *clusterProviderDescsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*logicalclusterv1alpha1.ClusterProviderDesc, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterProviderDescsResource, c.ClusterPath, name, pt, data, subresources...), &logicalclusterv1alpha1.ClusterProviderDesc{})
	if obj == nil {
		return nil, err
	}
	return obj.(*logicalclusterv1alpha1.ClusterProviderDesc), err
}
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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kubestellar/kubestellar/pkg/apis/edge/v1alpha1"
	scheme "github.com/kubestellar/kubestellar/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SyncTargetsGetter has a method to return a SyncTargetInterface.
// A group's client should implement this interface.
type SyncTargetsGetter interface {
	SyncTargets() SyncTargetInterface
}

// SyncTargetInterface has methods to work with SyncTarget resources.
type SyncTargetInterface interface {
	Create(ctx context.Context, syncTarget *v1alpha1.SyncTarget, opts v1.CreateOptions) (*v1alpha1.SyncTarget, error)
	Update(ctx context.Context, syncTarget *v1alpha1.SyncTarget, opts v1.UpdateOptions) (*v1alpha1.SyncTarget, error)
	UpdateStatus(ctx context.Context, syncTarget *v1alpha1.SyncTarget, opts v1.UpdateOptions) (*v1alpha1.SyncTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SyncTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SyncTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SyncTarget, err error)
	SyncTargetExpansion
}

// syncTargets implements SyncTargetInterface
type syncTargets struct {
	client rest.Interface
}

// newSyncTargets returns a SyncTargets
func newSyncTargets(c *EdgeV1alpha1Client) *syncTargets {
	return &syncTargets{
		client: c.RESTClient(),
	}
}

// Get takes name of the syncTarget, and returns the corresponding syncTarget object, and an error if there is any.
func (c *syncTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SyncTarget, err error) {
	result = &v1alpha1.SyncTarget{}
	err = c.client.Get().
		Resource("synctargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SyncTargets that match those selectors.
func (c *syncTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SyncTargetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SyncTargetList{}
	err = c.client.Get().
		Resource("synctargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested syncTargets.
func (c *syncTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("synctargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a syncTarget and creates it.  Returns the server's representation of the syncTarget, and an error, if there is any.
func (c *syncTargets) Create(ctx context.Context, syncTarget *v1alpha1.SyncTarget, opts v1.CreateOptions) (result *v1alpha1.SyncTarget, err error) {
	result = &v1alpha1.SyncTarget{}
	err = c.client.Post().
		Resource("synctargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(syncTarget).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a syncTarget and updates it. Returns the server's representation of the syncTarget, and an error, if there is any.
func (c *syncTargets) Update(ctx context.Context, syncTarget *v1alpha1.SyncTarget, opts v1.UpdateOptions) (result *v1alpha1.SyncTarget, err error) {
	result = &v1alpha1.SyncTarget{}
	err = c.client.Put().
		Resource("synctargets").
		Name(syncTarget.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(syncTarget).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *syncTargets) UpdateStatus(ctx context.Context, syncTarget *v1alpha1.SyncTarget, opts v1.UpdateOptions) (result *v1alpha1.SyncTarget, err error) {
	result = &v1alpha1.SyncTarget{}
	err = c.client.Put().
		Resource("synctargets").
		Name(syncTarget.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(syncTarget).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the syncTarget and deletes it. Returns an error if one occurs.
func (c *syncTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("synctargets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *syncTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("synctargets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched syncTarget.
func (c *syncTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SyncTarget, err error) {
	result = &v1alpha1.SyncTarget{}
	err = c.client.Patch(pt).
		Resource("synctargets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

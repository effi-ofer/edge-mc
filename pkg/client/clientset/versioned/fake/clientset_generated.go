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

package fake

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"

	clientset "github.com/kubestellar/kubestellar/pkg/client/clientset/versioned"
	edgev1alpha1 "github.com/kubestellar/kubestellar/pkg/client/clientset/versioned/typed/edge/v1alpha1"
	fakeedgev1alpha1 "github.com/kubestellar/kubestellar/pkg/client/clientset/versioned/typed/edge/v1alpha1/fake"
	metav1alpha1 "github.com/kubestellar/kubestellar/pkg/client/clientset/versioned/typed/meta/v1alpha1"
	fakemetav1alpha1 "github.com/kubestellar/kubestellar/pkg/client/clientset/versioned/typed/meta/v1alpha1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// EdgeV1alpha1 retrieves the EdgeV1alpha1Client
func (c *Clientset) EdgeV1alpha1() edgev1alpha1.EdgeV1alpha1Interface {
	return &fakeedgev1alpha1.FakeEdgeV1alpha1{Fake: &c.Fake}
}

// MetaV1alpha1 retrieves the MetaV1alpha1Client
func (c *Clientset) MetaV1alpha1() metav1alpha1.MetaV1alpha1Interface {
	return &fakemetav1alpha1.FakeMetaV1alpha1{Fake: &c.Fake}
}

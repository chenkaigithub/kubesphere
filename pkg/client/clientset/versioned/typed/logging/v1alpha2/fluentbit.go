/*
Copyright 2019 The KubeSphere authors.

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

package v1alpha2

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha2 "kubesphere.io/kubesphere/pkg/apis/logging/v1alpha2"
	scheme "kubesphere.io/kubesphere/pkg/client/clientset/versioned/scheme"
)

// FluentBitsGetter has a method to return a FluentBitInterface.
// A group's client should implement this interface.
type FluentBitsGetter interface {
	FluentBits(namespace string) FluentBitInterface
}

// FluentBitInterface has methods to work with FluentBit resources.
type FluentBitInterface interface {
	Create(*v1alpha2.FluentBit) (*v1alpha2.FluentBit, error)
	Update(*v1alpha2.FluentBit) (*v1alpha2.FluentBit, error)
	UpdateStatus(*v1alpha2.FluentBit) (*v1alpha2.FluentBit, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.FluentBit, error)
	List(opts v1.ListOptions) (*v1alpha2.FluentBitList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.FluentBit, err error)
	FluentBitExpansion
}

// fluentBits implements FluentBitInterface
type fluentBits struct {
	client rest.Interface
	ns     string
}

// newFluentBits returns a FluentBits
func newFluentBits(c *LoggingV1alpha2Client, namespace string) *fluentBits {
	return &fluentBits{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the fluentBit, and returns the corresponding fluentBit object, and an error if there is any.
func (c *fluentBits) Get(name string, options v1.GetOptions) (result *v1alpha2.FluentBit, err error) {
	result = &v1alpha2.FluentBit{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fluentbits").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FluentBits that match those selectors.
func (c *fluentBits) List(opts v1.ListOptions) (result *v1alpha2.FluentBitList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.FluentBitList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fluentbits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fluentBits.
func (c *fluentBits) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("fluentbits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a fluentBit and creates it.  Returns the server's representation of the fluentBit, and an error, if there is any.
func (c *fluentBits) Create(fluentBit *v1alpha2.FluentBit) (result *v1alpha2.FluentBit, err error) {
	result = &v1alpha2.FluentBit{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("fluentbits").
		Body(fluentBit).
		Do().
		Into(result)
	return
}

// Update takes the representation of a fluentBit and updates it. Returns the server's representation of the fluentBit, and an error, if there is any.
func (c *fluentBits) Update(fluentBit *v1alpha2.FluentBit) (result *v1alpha2.FluentBit, err error) {
	result = &v1alpha2.FluentBit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("fluentbits").
		Name(fluentBit.Name).
		Body(fluentBit).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *fluentBits) UpdateStatus(fluentBit *v1alpha2.FluentBit) (result *v1alpha2.FluentBit, err error) {
	result = &v1alpha2.FluentBit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("fluentbits").
		Name(fluentBit.Name).
		SubResource("status").
		Body(fluentBit).
		Do().
		Into(result)
	return
}

// Delete takes name of the fluentBit and deletes it. Returns an error if one occurs.
func (c *fluentBits) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fluentbits").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fluentBits) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fluentbits").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched fluentBit.
func (c *fluentBits) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.FluentBit, err error) {
	result = &v1alpha2.FluentBit{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("fluentbits").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

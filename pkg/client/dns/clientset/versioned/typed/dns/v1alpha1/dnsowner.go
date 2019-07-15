/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	"time"

	v1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	scheme "github.com/gardener/external-dns-management/pkg/client/dns/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DNSOwnersGetter has a method to return a DNSOwnerInterface.
// A group's client should implement this interface.
type DNSOwnersGetter interface {
	DNSOwners(namespace string) DNSOwnerInterface
}

// DNSOwnerInterface has methods to work with DNSOwner resources.
type DNSOwnerInterface interface {
	Create(*v1alpha1.DNSOwner) (*v1alpha1.DNSOwner, error)
	Update(*v1alpha1.DNSOwner) (*v1alpha1.DNSOwner, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DNSOwner, error)
	List(opts v1.ListOptions) (*v1alpha1.DNSOwnerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DNSOwner, err error)
	DNSOwnerExpansion
}

// dNSOwners implements DNSOwnerInterface
type dNSOwners struct {
	client rest.Interface
	ns     string
}

// newDNSOwners returns a DNSOwners
func newDNSOwners(c *DnsV1alpha1Client, namespace string) *dNSOwners {
	return &dNSOwners{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dNSOwner, and returns the corresponding dNSOwner object, and an error if there is any.
func (c *dNSOwners) Get(name string, options v1.GetOptions) (result *v1alpha1.DNSOwner, err error) {
	result = &v1alpha1.DNSOwner{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsowners").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DNSOwners that match those selectors.
func (c *dNSOwners) List(opts v1.ListOptions) (result *v1alpha1.DNSOwnerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DNSOwnerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsowners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dNSOwners.
func (c *dNSOwners) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dnsowners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dNSOwner and creates it.  Returns the server's representation of the dNSOwner, and an error, if there is any.
func (c *dNSOwners) Create(dNSOwner *v1alpha1.DNSOwner) (result *v1alpha1.DNSOwner, err error) {
	result = &v1alpha1.DNSOwner{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dnsowners").
		Body(dNSOwner).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dNSOwner and updates it. Returns the server's representation of the dNSOwner, and an error, if there is any.
func (c *dNSOwners) Update(dNSOwner *v1alpha1.DNSOwner) (result *v1alpha1.DNSOwner, err error) {
	result = &v1alpha1.DNSOwner{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnsowners").
		Name(dNSOwner.Name).
		Body(dNSOwner).
		Do().
		Into(result)
	return
}

// Delete takes name of the dNSOwner and deletes it. Returns an error if one occurs.
func (c *dNSOwners) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsowners").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dNSOwners) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsowners").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dNSOwner.
func (c *dNSOwners) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DNSOwner, err error) {
	result = &v1alpha1.DNSOwner{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dnsowners").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

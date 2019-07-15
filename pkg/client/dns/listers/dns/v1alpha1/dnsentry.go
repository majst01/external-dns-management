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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DNSEntryLister helps list DNSEntries.
type DNSEntryLister interface {
	// List lists all DNSEntries in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DNSEntry, err error)
	// DNSEntries returns an object that can list and get DNSEntries.
	DNSEntries(namespace string) DNSEntryNamespaceLister
	DNSEntryListerExpansion
}

// dNSEntryLister implements the DNSEntryLister interface.
type dNSEntryLister struct {
	indexer cache.Indexer
}

// NewDNSEntryLister returns a new DNSEntryLister.
func NewDNSEntryLister(indexer cache.Indexer) DNSEntryLister {
	return &dNSEntryLister{indexer: indexer}
}

// List lists all DNSEntries in the indexer.
func (s *dNSEntryLister) List(selector labels.Selector) (ret []*v1alpha1.DNSEntry, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DNSEntry))
	})
	return ret, err
}

// DNSEntries returns an object that can list and get DNSEntries.
func (s *dNSEntryLister) DNSEntries(namespace string) DNSEntryNamespaceLister {
	return dNSEntryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DNSEntryNamespaceLister helps list and get DNSEntries.
type DNSEntryNamespaceLister interface {
	// List lists all DNSEntries in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DNSEntry, err error)
	// Get retrieves the DNSEntry from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DNSEntry, error)
	DNSEntryNamespaceListerExpansion
}

// dNSEntryNamespaceLister implements the DNSEntryNamespaceLister
// interface.
type dNSEntryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DNSEntries in the indexer for a given namespace.
func (s dNSEntryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DNSEntry, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DNSEntry))
	})
	return ret, err
}

// Get retrieves the DNSEntry from the indexer for a given namespace and name.
func (s dNSEntryNamespaceLister) Get(name string) (*v1alpha1.DNSEntry, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dnsentry"), name)
	}
	return obj.(*v1alpha1.DNSEntry), nil
}

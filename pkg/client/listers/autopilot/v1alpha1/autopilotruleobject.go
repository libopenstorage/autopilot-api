/*
Copyright 2019 Openstorage.org

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
	v1alpha1 "github.com/libopenstorage/autopilot-api/pkg/apis/autopilot/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AutopilotRuleObjectLister helps list AutopilotRuleObjects.
type AutopilotRuleObjectLister interface {
	// List lists all AutopilotRuleObjects in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AutopilotRuleObject, err error)
	// AutopilotRuleObjects returns an object that can list and get AutopilotRuleObjects.
	AutopilotRuleObjects(namespace string) AutopilotRuleObjectNamespaceLister
	AutopilotRuleObjectListerExpansion
}

// autopilotRuleObjectLister implements the AutopilotRuleObjectLister interface.
type autopilotRuleObjectLister struct {
	indexer cache.Indexer
}

// NewAutopilotRuleObjectLister returns a new AutopilotRuleObjectLister.
func NewAutopilotRuleObjectLister(indexer cache.Indexer) AutopilotRuleObjectLister {
	return &autopilotRuleObjectLister{indexer: indexer}
}

// List lists all AutopilotRuleObjects in the indexer.
func (s *autopilotRuleObjectLister) List(selector labels.Selector) (ret []*v1alpha1.AutopilotRuleObject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutopilotRuleObject))
	})
	return ret, err
}

// AutopilotRuleObjects returns an object that can list and get AutopilotRuleObjects.
func (s *autopilotRuleObjectLister) AutopilotRuleObjects(namespace string) AutopilotRuleObjectNamespaceLister {
	return autopilotRuleObjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AutopilotRuleObjectNamespaceLister helps list and get AutopilotRuleObjects.
type AutopilotRuleObjectNamespaceLister interface {
	// List lists all AutopilotRuleObjects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AutopilotRuleObject, err error)
	// Get retrieves the AutopilotRuleObject from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AutopilotRuleObject, error)
	AutopilotRuleObjectNamespaceListerExpansion
}

// autopilotRuleObjectNamespaceLister implements the AutopilotRuleObjectNamespaceLister
// interface.
type autopilotRuleObjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AutopilotRuleObjects in the indexer for a given namespace.
func (s autopilotRuleObjectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AutopilotRuleObject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutopilotRuleObject))
	})
	return ret, err
}

// Get retrieves the AutopilotRuleObject from the indexer for a given namespace and name.
func (s autopilotRuleObjectNamespaceLister) Get(name string) (*v1alpha1.AutopilotRuleObject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("autopilotruleobject"), name)
	}
	return obj.(*v1alpha1.AutopilotRuleObject), nil
}

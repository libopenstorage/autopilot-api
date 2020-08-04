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

// ActionApprovalLister helps list ActionApprovals.
type ActionApprovalLister interface {
	// List lists all ActionApprovals in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ActionApproval, err error)
	// ActionApprovals returns an object that can list and get ActionApprovals.
	ActionApprovals(namespace string) ActionApprovalNamespaceLister
	ActionApprovalListerExpansion
}

// actionApprovalLister implements the ActionApprovalLister interface.
type actionApprovalLister struct {
	indexer cache.Indexer
}

// NewActionApprovalLister returns a new ActionApprovalLister.
func NewActionApprovalLister(indexer cache.Indexer) ActionApprovalLister {
	return &actionApprovalLister{indexer: indexer}
}

// List lists all ActionApprovals in the indexer.
func (s *actionApprovalLister) List(selector labels.Selector) (ret []*v1alpha1.ActionApproval, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ActionApproval))
	})
	return ret, err
}

// ActionApprovals returns an object that can list and get ActionApprovals.
func (s *actionApprovalLister) ActionApprovals(namespace string) ActionApprovalNamespaceLister {
	return actionApprovalNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ActionApprovalNamespaceLister helps list and get ActionApprovals.
type ActionApprovalNamespaceLister interface {
	// List lists all ActionApprovals in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ActionApproval, err error)
	// Get retrieves the ActionApproval from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ActionApproval, error)
	ActionApprovalNamespaceListerExpansion
}

// actionApprovalNamespaceLister implements the ActionApprovalNamespaceLister
// interface.
type actionApprovalNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ActionApprovals in the indexer for a given namespace.
func (s actionApprovalNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ActionApproval, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ActionApproval))
	})
	return ret, err
}

// Get retrieves the ActionApproval from the indexer for a given namespace and name.
func (s actionApprovalNamespaceLister) Get(name string) (*v1alpha1.ActionApproval, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("actionapproval"), name)
	}
	return obj.(*v1alpha1.ActionApproval), nil
}
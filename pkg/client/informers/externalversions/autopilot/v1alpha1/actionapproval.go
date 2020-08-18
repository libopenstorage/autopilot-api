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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	autopilotv1alpha1 "github.com/libopenstorage/autopilot-api/pkg/apis/autopilot/v1alpha1"
	versioned "github.com/libopenstorage/autopilot-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/libopenstorage/autopilot-api/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/libopenstorage/autopilot-api/pkg/client/listers/autopilot/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ActionApprovalInformer provides access to a shared informer and lister for
// ActionApprovals.
type ActionApprovalInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ActionApprovalLister
}

type actionApprovalInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewActionApprovalInformer constructs a new informer for ActionApproval type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewActionApprovalInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredActionApprovalInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredActionApprovalInformer constructs a new informer for ActionApproval type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredActionApprovalInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutopilotV1alpha1().ActionApprovals(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutopilotV1alpha1().ActionApprovals(namespace).Watch(options)
			},
		},
		&autopilotv1alpha1.ActionApproval{},
		resyncPeriod,
		indexers,
	)
}

func (f *actionApprovalInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredActionApprovalInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *actionApprovalInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&autopilotv1alpha1.ActionApproval{}, f.defaultInformer)
}

func (f *actionApprovalInformer) Lister() v1alpha1.ActionApprovalLister {
	return v1alpha1.NewActionApprovalLister(f.Informer().GetIndexer())
}

/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// This file was automatically generated by informer-gen

package internalversion

import (
	network "github.com/openshift/origin/pkg/network/apis/network"
	internalinterfaces "github.com/openshift/origin/pkg/network/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/network/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/network/generated/listers/network/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterNetworkInformer provides access to a shared informer and lister for
// ClusterNetworks.
type ClusterNetworkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ClusterNetworkLister
}

type clusterNetworkInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterNetworkInformer constructs a new informer for ClusterNetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterNetworkInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterNetworkInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterNetworkInformer constructs a new informer for ClusterNetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterNetworkInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Network().ClusterNetworks().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Network().ClusterNetworks().Watch(options)
			},
		},
		&network.ClusterNetwork{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterNetworkInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterNetworkInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterNetworkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&network.ClusterNetwork{}, f.defaultInformer)
}

func (f *clusterNetworkInformer) Lister() internalversion.ClusterNetworkLister {
	return internalversion.NewClusterNetworkLister(f.Informer().GetIndexer())
}

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	meteringv1 "github.com/kube-reporting/metering-operator/pkg/apis/metering/v1"
	versioned "github.com/kube-reporting/metering-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/kube-reporting/metering-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/kube-reporting/metering-operator/pkg/generated/listers/metering/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StorageLocationInformer provides access to a shared informer and lister for
// StorageLocations.
type StorageLocationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.StorageLocationLister
}

type storageLocationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewStorageLocationInformer constructs a new informer for StorageLocation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageLocationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageLocationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredStorageLocationInformer constructs a new informer for StorageLocation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageLocationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeteringV1().StorageLocations(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeteringV1().StorageLocations(namespace).Watch(options)
			},
		},
		&meteringv1.StorageLocation{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageLocationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageLocationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *storageLocationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&meteringv1.StorageLocation{}, f.defaultInformer)
}

func (f *storageLocationInformer) Lister() v1.StorageLocationLister {
	return v1.NewStorageLocationLister(f.Informer().GetIndexer())
}

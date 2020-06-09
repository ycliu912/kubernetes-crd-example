package main

import (
	"time"

	"github.com/ycliu912/kubernetes-crd-example/api/types/v1alpha1"
	client_v1alpha1 "github.com/ycliu912/kubernetes-crd-example/clientset/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

func WatchResoouces(clientSet client_v1alpha1.ExampleV1Alpha1Interface) cache.Store {
	projectStore, projectController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (retult runtime.Object, err error) {
				return clientSet.Projects("some-namespace").List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.Projects("some-namespace").Watch(lo)
			},
		},
		&v1alpha1.Project{},
		1*time.Second,
		cache.ResourceEventHandlerFuncs{},
	)

	go projectController.Run(wait.NeverStop)
	return projectStore
}

package main

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	appinformer "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	applisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type controller struct {
	clientSet      kubernetes.Interface
	depLister      applisters.DeploymentLister
	depCacheSynced cache.InformerSynced
	queue          workqueue.RateLimitingInterface
}

func newController(clientSet kubernetes.Interface, depInformer appinformer.DeploymentInformer) *controller {
	c := &controller{
		clientSet:      clientSet,
		depLister:      depInformer.Lister(),
		depCacheSynced: depInformer.Informer().HasSynced,
		queue:          workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ekspose"),
	}

	depInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    c.handleAdd,
			DeleteFunc: c.handleDel,
		},
	)
	return c
}

func (c *controller) run(ch <-chan struct{}) {
	fmt.Println("Staring Controller")
	if !cache.WaitForCacheSync(ch, c.depCacheSynced) {
		fmt.Println("Waiting for cache to be synced")
	}
	go wait.Until(c.worker, 1*time.Second, ch)
	<-ch
}
func (c *controller) worker() {
	for c.processItem() {

	}
}

func (c *controller) processItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	key,err:=cache.MetaNamespaceKeyFunc(item)
	if err!=nil{
		fmt.Printf("Getting key from Cache:%s\n",err.Error())
		return false
	}
	ns,name,err:=cache.SplitMetaNamespaceKey(key)
	if err!=nil{
		fmt.Println("Getting erro while spilting key")
		return false
	}
	err = c.syncDeployment(ns,name)
	if err!=nil{
		//re-try
		fmt.Println("Getting erro while spilting key")
		return false
	}
	return true
}
func (c *controller) syncDeployment(ns,name string) error{
	ctx:=context.Background()
	svc:=corev1.Service{}
	_,err:=c.clientSet.CoreV1().Services(ns).Create(ctx,&svc,metav1.CreateOptions{})
	if err != nil{
		fmt.Println("Error creating service")
		
	}

	return nil
}
func (c *controller) handleAdd(obj interface{}) {
	fmt.Println("Add called")
	c.queue.Add(obj)
}
func (c *controller) handleDel(obj interface{}) {
	fmt.Println("Del called")
}

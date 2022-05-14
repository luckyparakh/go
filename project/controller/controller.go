package main

import (
	"context"
	"fmt"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
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
	defer c.queue.Forget(item) // Once itme is processed succesfully remove from list, to avoid re-processing
	key, err := cache.MetaNamespaceKeyFunc(item)
	if err != nil {
		fmt.Printf("Getting key from Cache:%s\n", err.Error())
		return false
	}
	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		fmt.Println("Getting error while spilting key")
		return false
	}
	ctx := context.Background()
	_, err = c.clientSet.AppsV1().Deployments(ns).Get(ctx, name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		fmt.Printf("Handle delete of %v", name)
		if err := c.clientSet.CoreV1().Services(ns).Delete(ctx, name, metav1.DeleteOptions{}); err != nil {
			fmt.Printf("Error deleting svc:%s\n", err.Error())
		}
		if err := c.clientSet.NetworkingV1().Ingresses(ns).Delete(ctx, name, metav1.DeleteOptions{}); err != nil {
			fmt.Printf("Error deleting ingress:%s\n", err.Error())
		}
		return true
	}

	err = c.syncDeployment(ns, name)
	if err != nil {
		//re-try
		fmt.Println("Getting error while creating deployment")
		return false
	}
	return true
}
func (c *controller) syncDeployment(ns, name string) error {
	ctx := context.Background()
	dep, err := c.depLister.Deployments(ns).Get(name)
	if err != nil {
		fmt.Println("Error while getting deployment name")
	}
	svc := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dep.Name,
			Namespace: ns,
		},
		Spec: corev1.ServiceSpec{
			Selector: depPodLables(dep),
			Ports: []corev1.ServicePort{
				{
					Name: "http",
					Port: 80,
				},
			},
		},
	}
	_, err = c.clientSet.CoreV1().Services(ns).Create(ctx, &svc, metav1.CreateOptions{})

	if err != nil {
		fmt.Println("Error creating service")

	}

	return c.createIngress(ctx, svc)
}

func (c *controller) createIngress(ctx context.Context, svc corev1.Service) error {
	pathtype := "Prefix"
	ingress := netv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      svc.Name,
			Namespace: svc.Namespace,
			Annotations: map[string]string{
				"nginx.ingress.kubernetes.io/rewrite-target": "/",
			},
		},
		Spec: netv1.IngressSpec{
			Rules: []netv1.IngressRule{
				{
					IngressRuleValue: netv1.IngressRuleValue{
						HTTP: &netv1.HTTPIngressRuleValue{
							Paths: []netv1.HTTPIngressPath{
								{
									Path:     fmt.Sprintf("/%s", svc.Name),
									PathType: (*netv1.PathType)(&pathtype),
									Backend: netv1.IngressBackend{
										Service: &netv1.IngressServiceBackend{
											Name: svc.Name,
											Port: netv1.ServiceBackendPort{
												Number: 80,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	_, err := c.clientSet.NetworkingV1().Ingresses(svc.Namespace).Create(ctx, &ingress, metav1.CreateOptions{})
	return err
}
func depPodLables(dep *appsv1.Deployment) map[string]string {
	return dep.Spec.Template.Labels
}

func (c *controller) handleAdd(obj interface{}) {
	fmt.Println("Add called")
	c.queue.Add(obj)
}
func (c *controller) handleDel(obj interface{}) {
	fmt.Println("Del called")
	c.queue.Add(obj)
}

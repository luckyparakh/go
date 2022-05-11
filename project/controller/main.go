package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "./config", "Location of your kubeconfig file")
	cfg, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error %s building config.", err.Error())
		cfg, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("Error %s getting inCluster config.", err.Error())
		}
	}
	client, err := clientset.NewForConfig(cfg)
	if err != nil {
		fmt.Printf("Error %s loading configuration.", err.Error())
	}
	
	ctx := context.Background()
	pods,err:=client.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error %s fetching pods.", err.Error())
	}
	fmt.Println("Pods")
	for _, pod := range pods.Items {
		fmt.Printf("%s", pod.Name)
		fmt.Println()
	}

	ch := make(chan struct{})
	informers := informers.NewSharedInformerFactory(client, 10*time.Minute)
	c := newController(client, informers.Apps().V1().Deployments())
	informers.Start(ch)
	c.run(ch)
	fmt.Println(informers)
}

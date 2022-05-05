package main

import (
	"context"
	"fmt"
	"flag"

	clientset "k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Keeping it over here because not able to parse path in Windows style
	kubeconfig:=flag.String("kubeconfig","./config","Location of your kubeconfig file")
	cfg, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Error %s building config.",err.Error())
	}
	client := clientset.NewForConfigOrDie(cfg)
	ctx:=context.Background()
	pods,err:=client.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error %s fetching pods.",err.Error())
	}
	fmt.Println("Pods")
	for _, pod := range pods.Items{
		fmt.Printf("%s", pod.Name)
	}
	deployments,err := client.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error %s fetching deployments.",err.Error())
	}
	fmt.Println("Deployments")
	for _, deployment := range deployments.Items{
		fmt.Printf("%s", deployment.Name)
	}
}
package main

import (
	"context"
	"fmt"

	"github.com/openshift/hive/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	fmt.Println("hello")

	kubeconfig := "/home/apodhrad/Projects/integration-qe/jcasc-declaration/secrets/hive-kubeconfig"

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	listNs, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	for _, ns := range listNs.Items {
		fmt.Println(ns.Name)
	}

	client, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	list, err := client.HiveV1().ClusterPools("integration-qe").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, clusterPool := range list.Items {
		fmt.Printf("%v\n", clusterPool.Name)
	}
}

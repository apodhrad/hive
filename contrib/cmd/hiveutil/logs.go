package main

import (
	"context"
	"fmt"

	"github.com/openshift/hive/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	pool      string
	namespace string
)

func LogsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs",
		Short: "Prints logs for a pod in a namespace associated to a given pool",
		Long: `Prints logs for a pod in a namespace associated to a given pool.
If there are more such namespaces then you need to specify it with -n or --namespace.`,
		Run: func(cmd *cobra.Command, args []string) {
			getNamespaces()
		},
	}

	pFlags := cmd.PersistentFlags()
	pFlags.StringVarP(&pool, "pool", "p", "", "Cluster pool name")
	pFlags.StringVarP(&namespace, "namespace", "n", "", "Openshift namespace")
	return cmd
}

func getNamespaces() {
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
	listNs, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{LabelSelector: "hive.openshift.io/cluster-pool-name=apodhrad-001"})
	for _, ns := range listNs.Items {
		fmt.Println(ns.Name)
		listPods, err := clientset.CoreV1().Pods(ns.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for _, pod := range listPods.Items {
			fmt.Println(pod.Name)
			req := clientset.CoreV1().Pods(ns.Name).GetLogs(pod.Name, &v1.PodLogOptions{})
			req.Stream()
		}
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

package main

import (
	"context"
	"flag"
	"log"

	corev1 "k8s.io/api/core/v1"

	k8srestclient "github.com/gufranmirza/k8s-controller-runtime-client/k8s-rest-client"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var kubeconfig string

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.Parse()
}

func main() {
	config, err := k8srestclient.LoadKubeConfig(kubeconfig)
	if err != nil {
		log.Fatalf("config load err %v\n", err)
	}

	c, err := k8srestclient.NewK8sRestClient(config, client.Options{})
	if err != nil {
		log.Fatalf("failed to create client with err %v\n", err)
	}

	pods := &corev1.PodList{}

	err = c.List(context.Background(), pods, client.InNamespace("default"))
	if err != nil {
		log.Fatalf("failed to list pods in namespace default with err %v\n", err)
	}

	for _, pod := range pods.Items {
		log.Println(pod.Name)
	}
}

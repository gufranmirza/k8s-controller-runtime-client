package k8srestclient

import (
	"fmt"
	"log"

	// cachev1alpha1 "memcached-operator/pkg/apis/cache/v1alpha1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// K8sRestClient defines a contract for kube rest clients
//go:generate counterfeiter . K8sRestClient
type K8sRestClient client.Client

func NewK8sRestClient(config *rest.Config, options client.Options) (K8sRestClient, error) {
	rc, err := client.New(config, options)
	if err != nil {
		return nil, fmt.Errorf("failed to create new kube rest client with error %v", err)
	}

	return rc, err
}

// BuildScheme returns an API scheme aggregated from resource definitions. Should
// not be invoked by other packages.
func BuildScheme() (*runtime.Scheme, error) {
	runtimeScheme := runtime.NewScheme()

	// add core go client scheme
	err := scheme.AddToScheme(runtimeScheme)
	if err != nil {
		return nil, fmt.Errorf("failed to add core schemes with error %v", err)
	}

	// err = cachev1alpha1.AddToScheme(runtimeScheme)
	// if err != nil {
	// 	return nil, errors.Wrap(
	// 		err,
	// 		"Failed to add operator scheme",
	// 	)
	// }

	return runtimeScheme, nil
}

func LoadKubeConfig(kubeconfig string) (*rest.Config, error) {
	var config *rest.Config
	var err error

	if kubeconfig == "" {
		log.Printf("using in-cluster configuration")
		config, err = rest.InClusterConfig()
	} else {
		log.Printf("using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to read kube config with error %v", err)
	}

	return config, nil
}

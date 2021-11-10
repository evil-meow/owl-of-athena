package k8s_api

import (
	"context"
	"log"

	"github.com/pytimer/k8sutil/apply"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func Apply(crdContents string) error {

	_, dynamicClient, discoveryClient, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	applyOptions := apply.NewApplyOptions(dynamicClient, discoveryClient)
	if err := applyOptions.Apply(context.TODO(), []byte(crdContents)); err != nil {
		log.Fatalf("apply error: %v", err)
	}

	return nil
}

func CopySecret(sourceName string, sourceNamespace string, targetNamespace string) error {
	clientset, _, _, err := connect()
	if err != nil {
		return err
	}
	secret, err := clientset.CoreV1().Secrets(sourceNamespace).Get(context.TODO(), sourceName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	secret.ObjectMeta.Namespace = targetNamespace
	secret.ResourceVersion = ""

	_, err = clientset.CoreV1().Secrets(targetNamespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func connect() (*kubernetes.Clientset, dynamic.Interface, *discovery.DiscoveryClient, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, nil, nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, err
	}

	dynamic, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, err
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset, dynamic, discoveryClient, nil
}

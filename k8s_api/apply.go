package k8s_api

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func Apply(url string) error {
	_, err := connect()
	if err != nil {
		return err
	}
	return nil
}

func CopySecret(sourceName string, sourceNamespace string, targetNamespace string) error {
	clientset, err := connect()
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

func connect() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

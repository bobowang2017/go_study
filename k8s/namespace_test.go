package k8s

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestListNameSpace(t *testing.T) {
	clientset := GetClient()
	ns, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for idx, _ns := range ns.Items {
		fmt.Println(idx, _ns.Name)
	}
}

func TestCreateNameSpace(t *testing.T) {
	clientset := GetClient()
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "helloworld",
		},
	}
	ns, err := clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(ns)
}

func TestDeleteNameSpace(t *testing.T) {
	clientset := GetClient()
	err := clientset.CoreV1().Namespaces().Delete(context.TODO(), "helloworld", metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}
}

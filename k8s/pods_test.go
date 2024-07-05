package k8s

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestListPods(t *testing.T) {
	clientset := GetClient()
	pods, err := clientset.CoreV1().Pods("ccpow-zqdev01").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for idx, pod := range pods.Items {
		fmt.Println(idx, pod.Name)
	}
}

//func TestCreateNameSpace(t *testing.T) {
//	clientset := GetClient()
//	namespace := &corev1.Namespace{
//		ObjectMeta: metav1.ObjectMeta{
//			Name: "helloworld",
//		},
//	}
//	ns, err := clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
//	if err != nil {
//		panic(err.Error())
//	}
//	fmt.Println(ns)
//}
//
//func TestDeleteNameSpace(t *testing.T) {
//	clientset := GetClient()
//	err := clientset.CoreV1().Namespaces().Delete(context.TODO(), "helloworld", metav1.DeleteOptions{})
//	if err != nil {
//		panic(err.Error())
//	}
//}

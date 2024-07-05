package k8s

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"testing"
)

func TestListNode(t *testing.T) {
	clientset := GetClient()
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for idx, node := range nodes.Items {
		fmt.Println(idx, node.Name)
	}
}

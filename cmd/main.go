package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type NetworkConfiguration struct {
	XMLName      xml.Name `xml:"NetworkConfiguration"`
	KnownProxies []string `xml:"KnownProxies>string"`
}

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		ips := []string{}
		namespace := os.Getenv("NAMESPACE")
		labelSelector := os.Getenv("LABEL_SELECTOR")
		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
			LabelSelector: labelSelector,
		})
		if err != nil {
			fmt.Println("Error fetching pods:", err)
			continue
		}

		for _, pod := range pods.Items {
			if pod.Status.PodIP != "" {
				ips = append(ips, pod.Status.PodIP)
			}
		}

		config := NetworkConfiguration{KnownProxies: ips}
		file, err := os.Create("/config/network.xml")
		if err != nil {
			fmt.Println("Error writing file:", err)
			continue
		}
		defer file.Close()

		encoder := xml.NewEncoder(file)
		encoder.Indent("", "  ")
		if err := encoder.Encode(config); err != nil {
			fmt.Println("Error encoding XML:", err)
		} else {
			fmt.Println("Updated network.xml with IPs:", ips)
		}

		time.Sleep(300 * time.Second)
	}
}

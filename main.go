package main

import (
	"fmt"
	"time"

	helmv1 "github.com/fluxcd/helm-operator/pkg/apis/helm.fluxcd.io/v1"
	helmclient "github.com/fluxcd/helm-operator/pkg/client/clientset/versioned/typed/helm.fluxcd.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func main() {

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("Failed to create config: %v", err.Error())
		panic(err.Error())
	}

	// creates the helmclient
	hc, err := helmclient.NewForConfig(config)
	if err != nil {
		fmt.Printf("Failed to create helmclient: %v", err.Error())
		panic(err.Error())
	}

	createHelmRelease(config, hc, "test2")
	createHelmRelease(config, hc, "test3")
	for {
		listHelmReleases(config, hc, "test")
		time.Sleep(10 * time.Second)
	}

}

func listHelmReleases(c *rest.Config, hc *helmclient.HelmV1Client, ns string) {
	lo := metav1.ListOptions{
		LabelSelector: "konghq.com/provisioner-managed",
	}
	hrl, err := hc.HelmReleases(ns).List(lo)
	if err != nil {
		fmt.Printf("Failed to create HelmRelease: %v\n", err.Error())
	} else {
		fmt.Printf("Found %d releases:\n", len(hrl.Items))
		for _, hr := range hrl.Items {
			fmt.Printf("Name: %s, Namespace: %s\n", hr.Name, hr.Namespace)
		}

	}
}

func createHelmRelease(c *rest.Config, hc *helmclient.HelmV1Client, name string) {
	////////// create a test HelmRelease
	//	---
	//	apiVersion: helm.fluxcd.io/v1
	//	kind: HelmRelease
	//	metadata:
	//		name: metrics-server
	//	spec:
	//		releaseName: metrics-server
	//		chart:
	//			repository: https://charts.bitnami.com/bitnami
	//			name: metrics-server
	//			version: 4.5.3
	//		values:
	//			apiService:
	//				create: true

	rcs := helmv1.RepoChartSource{
		RepoURL: "https://charts.test.com",
		Name:    name,
		Version: "1.0.0",
	}

	hr := helmv1.HelmRelease{
		TypeMeta: metav1.TypeMeta{
			Kind:       "HelmRelease",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "test",
			Labels: map[string]string{
				"konghq.com/provisioner-managed": "true",
			},
		},
		Spec: helmv1.HelmReleaseSpec{
			ReleaseName: name,
			ChartSource: helmv1.ChartSource{
				RepoChartSource: &rcs,
			},
		},
	}

	_, err := hc.HelmReleases("test").Create(&hr)
	if err != nil {
		fmt.Printf("Failed to create HelmRelease: %v\n", err.Error())
	} else {
		fmt.Println("seems to have worked :P")
	}

}

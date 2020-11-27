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
	fmt.Println("hello, this app is working...")

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Errorf("Failed 49: %v", err.Error())
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := helmclient.NewForConfig(config)
	if err != nil {
		fmt.Errorf("Failed 59: %v", err.Error())
		panic(err.Error())
	}

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
		RepoURL: "https://charts.bitnami.com/bitnami",
		Name:    "metrics-server",
		Version: "4.5.3",
	}

	hr := helmv1.HelmRelease{
		TypeMeta: metav1.TypeMeta{
			Kind:       "HelmRelease",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test1",
			Namespace: "test",
		},
		Spec: helmv1.HelmReleaseSpec{
			ReleaseName: "test",
			ChartSource: helmv1.ChartSource{
				RepoChartSource: &rcs,
			},
		},
	}

	for {
		_, err = clientset.HelmReleases("test").Create(&hr)
		if err != nil {
			fmt.Println("seems to have failed :(")
			fmt.Printf("Failed 113: %v\n", err.Error())
		} else {
			fmt.Println("seems to have worked :P")
		}
		time.Sleep(10 * time.Second)
	}

}

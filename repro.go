package main

import (
	"context"
	"flag"
	"path/filepath"

	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var configPath string
	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(&configPath, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		flag.StringVar(&configPath, "kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		panic(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pods := client.CoreV1().Pods("ns")

	pods.GetLogs("foo", nil) // <- this is default Ctrl+Space autocomplete result for "opts" arg

	pods.Create(context.Background(), &core.Pod{}, meta.CreateOptions{})
	//                                             ^ here
	// autocomplete (Cltr+Shift+Space) for third argument (opts metav1.CreateOption) is pretty slow and
	// works only if `meta "k8s.io/apimachinery/pkg/apis/meta/v1"` package is imported.
	pods.EvictV1(context.Background(), nil)
	//                                 ^ here
	// Ctrl+Shift+Space produces "nil", probably because IDE can't find 'policyv1 "k8s.io/api/policy/v1" package.
}

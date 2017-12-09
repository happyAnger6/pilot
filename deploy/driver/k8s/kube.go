package k8s

import (
	"os"
	"pilot/deploy/driver"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/core/v1"
	"path/filepath"
	"fmt"
)

const (
	driverName = "kube"
)

type Driver struct{
	clientSet *kubernetes.Clientset
}

func (d *Driver) String() string {
	return driverName
}

func (d *Driver) StartContainer(name string, opts *driver.ContainerOpts) error {
	containers := make([]v1.Container, 1)
	containers[0] = v1.Container{
		Name:"v9simware",
		Image:"comware.io/v9mpu:trunk",
	}
	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name + opts.CreateOpts["bchassis"] + opts.CreateOpts["bslot"] + opts.CreateOpts["bcpu"],
			Annotations: map[string]string{"network_info":`aa`},
		},
		Spec: v1.PodSpec{
			Containers: containers,
			NodeName:   opts.CreateOpts["brunnode"],
		},
	}
	d.clientSet.CoreV1().Pods("default").Create(pod)
	return nil
}

func (d *Driver) StopContainer(name string) error {
	return nil
}

func (d *Driver) RemoveContainer(name string) error {
	return nil
}

// use the current context in kubeconfig
func Init() (driver.Driver, error){
	var kubeConfigPath string
	if home := homeDir(); home != ""{
		kubeConfigPath = filepath.Join(home, ".kube", "config")
	}
	fmt.Printf("use config file:%v\r\n", kubeConfigPath)

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	d := &Driver{
		clientSet:	clientset,
	}

	return d, nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
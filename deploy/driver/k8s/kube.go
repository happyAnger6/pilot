package k8s

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/core/v1"
	log "github.com/sirupsen/logrus"

	"pilot/deploy/driver"
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
	log.Debugf("StartContainer: %v\r\n", opts)

	chassis := opts.CreateOpts["bchassis"].(string)
	slot := opts.CreateOpts["bslot"].(string)
	cpu := opts.CreateOpts["bcpu"].(string)
	projName := opts.CreateOpts["bname"].(string)
	bType := opts.CreateOpts["btype"].(string)
	selfNode := chassis + "," + slot + "," + cpu
	sep := "-"
	podName := projName + sep + bType + sep + chassis + sep + slot + sep + cpu
	privileged := true

	containers := make([]v1.Container, 1)
	containers[0] = v1.Container{
		Name: "v9simware",
		Image: opts.CreateOpts["bimage"].(string),
		Stdin: true,
		TTY: true,
		SecurityContext: &v1.SecurityContext{Privileged: &privileged},
		Env: []v1.EnvVar{{Name: "SELFNODE", Value: selfNode}},
		VolumeMounts: []v1.VolumeMount{{Name:"drvbm", MountPath:"/var/drv/bm"}},
		Command: []string{"/bin/bash"},
		Args: []string{"/bin/v9.sh"},
	}

	hostPath := &v1.HostPathVolumeSource{Path: "/var/drv/bm"}
	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
			Annotations: map[string]string{"network_info":`aa`},
		},
		Spec: v1.PodSpec{
			Containers: containers,
			NodeName:   opts.CreateOpts["brunnode"].(string),
			Volumes: []v1.Volume{{Name: "drvbm", VolumeSource: v1.VolumeSource{HostPath: hostPath}}},
		},
	}
	pod, err := d.clientSet.CoreV1().Pods("default").Create(pod)
	return err
}

func (d *Driver) ListContainers() (*driver.ContainerList, error) {
	lists := driver.ContainerList{}
	pods, err := d.clientSet.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _ , pod := range pods.Items {
		lists.Items = append(lists.Items, pod)
	}
	return &lists, nil
}

func (d *Driver) StopContainer(name string) error {
	err := d.clientSet.CoreV1().Pods("default").Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (d *Driver) RemoveContainer(name string) error {
	err := d.clientSet.CoreV1().Pods("default").Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

// use the current context in kubeconfig
func Init() (driver.Driver, error){
	var kubeConfigPath string
	if home := homeDir(); home != ""{
		kubeConfigPath = filepath.Join(home, ".kube", "config")
	}
	log.Debugf("use config file:%v\r\n", kubeConfigPath)

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
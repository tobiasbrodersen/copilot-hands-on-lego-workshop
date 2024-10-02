package main

import (
    "context"
    "flag"
    "fmt"
    "os"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/util/homedir"
    "path/filepath"
    "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/util/intstr"
    "k8s.io/api/apps/v1beta1"
)

func scaleDeployment(namespace, deploymentName string, replicas int32) {
    // Load kube config from default location
    var kubeconfig *string
    if home := homedir.HomeDir(); home != "" {
        kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    } else {
        kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    flag.Parse()

    // Build the config from the kubeconfig file
    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        fmt.Printf("Error building kubeconfig: %v\n", err)
        os.Exit(1)
    }

    // Create the clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        fmt.Printf("Error creating clientset: %v\n", err)
        os.Exit(1)
    }

    // Define the scale object
    scale := &v1beta1.Scale{
        Spec: v1beta1.ScaleSpec{
            Replicas: replicas,
        },
    }

    // Scale the deployment
    _, err = clientset.AppsV1beta1().Deployments(namespace).UpdateScale(context.TODO(), deploymentName, scale, v1.UpdateOptions{})
    if err != nil {
        fmt.Printf("Error scaling deployment: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("Scaled deployment %s to %d replicas\n", deploymentName, replicas)
}

func main() {
    namespace := "default"  // Replace with your namespace if different
    deploymentName := "myapp-deployment"
    replicas := int32(5)

    scaleDeployment(namespace, deploymentName, replicas)
}

from kubernetes import client, config

def scale_deployment(namespace, deployment_name, replicas):
    # Load kube config from default location
    config.load_kube_config()

    # Create an instance of the AppsV1Api
    api_instance = client.AppsV1Api()

    # Define the scale object
    scale = client.V1Scale(
        spec=client.V1ScaleSpec(
            replicas=replicas
        )
    )

    # Scale the deployment
    api_instance.replace_namespaced_deployment_scale(
        name=deployment_name,
        namespace=namespace,
        body=scale
    )

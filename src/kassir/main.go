package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"os"
)

func main() {
	endpoint := os.Getenv("DOCKER_HOST")
	path := os.Getenv("DOCKER_CERT_PATH")
	ca := fmt.Sprintf("%s/ca.pem", path)
	cert := fmt.Sprintf("%s/cert.pem", path)
	key := fmt.Sprintf("%s/key.pem", path)
	client, _ := docker.NewTLSClient(endpoint, cert, key, ca)
	containers, _ := client.ListContainers(docker.ListContainersOptions{})
	for _, cnt := range containers {
		fmt.Println("ID: ", cnt.ID)
		fmt.Println("RepoTags: ", cnt.Labels)
	}
}

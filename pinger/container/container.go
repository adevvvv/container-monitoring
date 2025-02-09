package container

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetContainerIPs() ([]string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("error connecting to Docker API: %v", err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error getting container list: %v", err)
	}

	var ips []string
	for _, container := range containers {
		inspect, err := cli.ContainerInspect(context.Background(), container.ID)
		if err != nil {
			log.Printf("Error getting container info %s: %v", container.ID, err)
			continue
		}
		for _, netConf := range inspect.NetworkSettings.Networks {
			ips = append(ips, netConf.IPAddress)
		}
	}
	return ips, nil
}

func PingContainer(ip string) (float64, error) {
	start := time.Now()
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	err := cmd.Run()
	if err != nil {
		return 0, fmt.Errorf("container %s did not respond to ping", ip)
	}
	elapsed := time.Since(start).Seconds() * 1000
	return elapsed, nil
}

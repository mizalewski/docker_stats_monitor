package docker_stats_monitor

import (
	"github.com/docker/docker/client"
	"io/ioutil"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"context"
	"github.com/docker/docker/daemon/stats"
	"io"
)

type DockerApiClient struct {
	cli *client.Client
}

func NewApiClient() (*DockerApiClient, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	apiClient := &DockerApiClient{cli: cli}
	return apiClient, nil
}

func (apiClient *DockerApiClient) GetContainersStats() ([]ContainerStats, error) {
	cli := apiClient.cli

	ctx := context.Background()
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	for _, container := range containers {
		getContainerStats(cli, ctx, container)
	}
}

func getContainerStats(cli *client.Client, ctx context.Context, container types.Container) (*ContainerStats, error) {
	response, err := cli.ContainerStats(ctx, container.ID, false)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	containerStats, err := parseContainerStatsResponse(response.Body)
	if err != nil {
		return nil, err
	}

	return containerStats, nil
}

func parseContainerStatsResponse(response io.ReadCloser) (*ContainerStats, error) {
	body, err := ioutil.ReadAll(response)
	if err != nil {
		return nil, err
	}

	containerStats := &ContainerStats{}
	err = json.Unmarshal(body, containerStats)
	if err != nil {
		return nil, err
	}
}

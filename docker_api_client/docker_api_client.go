package docker_api_client

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"io/ioutil"
	"strings"
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

	var containersStats []ContainerStats
	for _, container := range containers {
		stats, err := getContainerStats(cli, ctx, container)
		if err != nil {
			return nil, err
		}

		image, imageTag := extractImageAndTag(container.Image)
		stats.Image = image
		stats.ImageTag = imageTag
		containersStats = append(containersStats, *stats)
	}

	return containersStats, nil
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

	return containerStats, nil
}

func extractImageAndTag(image string) (string, string) {
	splitted := strings.Split(image, ":")
	if len(splitted) > 1 {
		return splitted[0], strings.Join(splitted[1:], ":")
	}

	return splitted[0], "latest"
}

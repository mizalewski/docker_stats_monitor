package main

import (
	"encoding/json"
	"fmt"
	"github.com/mizalewski/docker_stats_monitor/docker_api_client"
	"time"
)

const (
	SleepBetweenReports = 30 * time.Second
)

func main() {
	for {
		reportDockerStats()
		time.Sleep(SleepBetweenReports)
	}
}

func reportDockerStats() {
	dockerApiClient, err := docker_api_client.NewApiClient()
	if err != nil {
		panic(err)
	}
	containersStats, err := dockerApiClient.GetContainersStats()
	if err != nil {
		panic(err)
	}

	for _, stats := range containersStats {
		printJsonFormattedStats(&stats)
	}
}

func printJsonFormattedStats(stats *docker_api_client.ContainerStats) {
	jsonFormatted, err := json.Marshal(stats)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonFormatted))
}

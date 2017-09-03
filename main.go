package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/mizalewski/docker_stats_monitor/aws_metrics"
	"github.com/mizalewski/docker_stats_monitor/docker_api_client"
	"os"
	"time"
)

const (
	SleepBetweenReports = 30 * time.Second
)

func main() {
	dockerApiClient := createDockerApiClient()
	awsMetricsClient := createAwsMetricsClient(os.Getenv("METRICS_NAMESPACE"))

	for {
		containerStats := readDockerStats(dockerApiClient)
		sendStatsToCloudWatch(awsMetricsClient, containerStats)

		time.Sleep(SleepBetweenReports)
	}
}
func sendStatsToCloudWatch(awsMetricsClient *aws_metrics.AwsMetricsClient, stats []docker_api_client.ContainerStats) {
	metricsList := []*cloudwatch.MetricDatum{}

	for _, containerStats := range stats {
		metricsList = append(metricsList, mapStatsToMetrics(containerStats))
	}

	err := awsMetricsClient.SendMetrics(metricsList)
	if err != nil {
		panic(err)
	}
}
func mapStatsToMetrics(containerStats docker_api_client.ContainerStats) *cloudwatch.MetricDatum {
	return &cloudwatch.MetricDatum{
		MetricName: aws.String("Memory usage"),
		Unit:       aws.String(cloudwatch.StandardUnitBytes),
		Value:      aws.Float64(float64(containerStats.MemoryStats.Usage)),
		Dimensions: []*cloudwatch.Dimension{
			&cloudwatch.Dimension{
				Name:  aws.String("Name"),
				Value: aws.String(containerStats.Name),
			},
			&cloudwatch.Dimension{
				Name:  aws.String("Image"),
				Value: aws.String(containerStats.Image),
			},
		},
	}
}

func createDockerApiClient() *docker_api_client.DockerApiClient {
	dockerApiClient, err := docker_api_client.NewApiClient()
	if err != nil {
		panic(err)
	}
	return dockerApiClient
}

func readDockerStats(dockerApiClient *docker_api_client.DockerApiClient) []docker_api_client.ContainerStats {
	containersStats, err := dockerApiClient.GetContainersStats()
	if err != nil {
		panic(err)
	}

	return containersStats
}

func createAwsMetricsClient(metricsNamespace string) *aws_metrics.AwsMetricsClient {
	awsMetricsClient, err := aws_metrics.NewMetricsClient(metricsNamespace)
	if err != nil {
		panic(err)
	}

	return awsMetricsClient
}

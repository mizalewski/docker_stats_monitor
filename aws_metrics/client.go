package aws_metrics

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

type AwsMetricsClient struct {
	cloudwatchService *cloudwatch.CloudWatch
	namespace         string
}

func NewMetricsClient(namespace string) (*AwsMetricsClient, error) {
	//awsSession, err := session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: session.SharedConfigEnable,
	//})
	awsSession, err := session.NewSession()

	if err != nil {
		return nil, err
	}

	service := cloudwatch.New(awsSession)
	metricsClient := AwsMetricsClient{cloudwatchService: service, namespace: namespace}

	return &metricsClient, nil
}

func (metricsClient *AwsMetricsClient) SendMetrics(metricsData []*cloudwatch.MetricDatum) error {
	_, err := metricsClient.cloudwatchService.PutMetricData(&cloudwatch.PutMetricDataInput{
		MetricData: metricsData,
		Namespace:  aws.String(metricsClient.namespace),
	})

	return err
}

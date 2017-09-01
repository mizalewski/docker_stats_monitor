package docker_stats_monitor

import (
	"github.com/docker/docker/client"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/mizalewski/docker_stats_monitor/docker_api_client"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func main() {
	dockerApiClient, err := docker_api_client.NewApiClient()
	if err != nil {
		panic(err)
	}
}

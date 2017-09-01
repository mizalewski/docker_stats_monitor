package docker_api_client

type ContainerStats struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Image    string
	ImageTag string

	CpuStats struct {
		CpuUsage struct {
			TotalUsage int64 `json:"total_usage"`
		} `json:"cpu_usage"`
	} `json:"cpu_stats"`

	MemoryStats struct {
		Usage    int64 `json:"usage"`
		MaxUsage int64 `json:"max_usage"`
	} `json:"memory_stats"`
}

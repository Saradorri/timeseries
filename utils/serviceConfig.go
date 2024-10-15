package utils

type ServiceConfig struct {
	App      app      `mapstructure:"app"`
	Database database `mapstructure:"database"`
}

type app struct {
	GrpcPort               int    `mapstructure:"grpc_port"`
	MetricPort             int    `mapstructure:"metric_port"`
	ScheduleIntervalMinute int    `mapstructure:"schedule_interval_min"`
	ApiUrl                 string `mapstructure:"api_url"`
}

type database struct {
	ScrapeURL string `mapstructure:"scrape_url"`
	JobName   string `mapstructure:"job_name"`
}

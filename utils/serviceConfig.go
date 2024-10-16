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
	InfluxDB influxdb `mapstructure:"influxdb"`
}

type influxdb struct {
	Url    string `mapstructure:"url"`
	Token  string `mapstructure:"token"`
	Org    string `mapstructure:"org"`
	Bucket string `mapstructure:"bucket"`
}

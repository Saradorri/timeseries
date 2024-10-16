package influxdb

import (
	"github.com/influxdata/influxdb-client-go/v2"
)

type InfluxDBClient struct {
	Client influxdb2.Client
}

func NewClient(url, token string) *InfluxDBClient {
	client := influxdb2.NewClient(url, token)
	return &InfluxDBClient{
		Client: client,
	}
}

func (c *InfluxDBClient) Close() {
	c.Client.Close()
}

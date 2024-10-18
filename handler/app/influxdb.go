package app

import (
	"edgecom.ai/timeseries/internal/repository"
	"edgecom.ai/timeseries/internal/repository/influxdb"
)

func (a *application) InitInfluxDB() repository.Repository {
	client := influxdb.NewClient(
		a.config.Database.InfluxDB.Url,
		a.config.Database.InfluxDB.Token,
	)
	return influxdb.NewRepository(client, a.config.Database.InfluxDB.Org, a.config.Database.InfluxDB.Bucket)
}

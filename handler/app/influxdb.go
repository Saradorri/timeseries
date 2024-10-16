package app

import (
	"edgecom.ai/timeseries/internal/repository/influxdb"
)

func (a *application) InitInfluxDB() influxdb.Repository {
	client := influxdb.NewClient(
		a.config.Database.InfluxDB.Url,
		a.config.Database.InfluxDB.Token,
	)
	return influxdb.NewRepository(client, a.config.Database.InfluxDB.Org, a.config.Database.InfluxDB.Bucket)
}

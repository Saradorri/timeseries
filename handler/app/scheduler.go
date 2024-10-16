package app

import (
	"edgecom.ai/timeseries/internal/scheduler"
	"edgecom.ai/timeseries/internal/services"
)

func (a *application) InitScheduler(service services.TimeSeriesScraperService) scheduler.Scheduler {
	return scheduler.NewScheduler(a.config.App.ApiUrl, a.config.App.ScheduleIntervalMinute, service)
}

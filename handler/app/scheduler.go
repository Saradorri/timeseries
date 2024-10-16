package app

import (
	"edgecom.ai/timeseries/internal/scheduler"
	"edgecom.ai/timeseries/internal/services"
)

func (a *application) InitScheduler(scraper services.ScraperService, ts services.TimeSeriesService) scheduler.Scheduler {
	return scheduler.NewScheduler(a.config.App.ScheduleIntervalMinute, scraper, ts)
}

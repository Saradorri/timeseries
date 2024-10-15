package main

import (
	"context"
	"edgecom.ai/timeseries/handler/app"
)

func main() {
	application := app.NewApplication(context.Background())
	application.Setup()
}

package main

import (
	"github.com/kirillVladov/s3-debugger/cmd/latency/app"
)

func main() {
	newApp, err := app.NewApp()

	if err != nil {
		panic(err)
	}

	newApp.Run()
}

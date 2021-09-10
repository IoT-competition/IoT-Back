package main

import (
	"github.com/IoTCompetition/IoBack/app"
)

func main() {
	server := app.BuildInjector()

	_ = server.ListenAndServe()
}

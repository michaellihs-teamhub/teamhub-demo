package main

import (
	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
	"github.com/michaellihs/gorilla-testing/app"
	"os"
)

func main() {
	initLogging()

	gorillaTestingApp := initApp()

	if err := gorillaTestingApp.Run(); err != nil {
		log.Error(err)
	}
}

func initApp() *app.GorillaTestingApp {
	port := os.Getenv("GORILLA_TESTING_PORT")
	gorillaTestingApp := app.NewGorillaTestingApp(port)
	return gorillaTestingApp
}

func initLogging() {
	cLog := console.New(true)
	log.AddHandler(cLog, log.AllLevels...)
}
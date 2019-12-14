package main

import (
	"github.com/michaellihs/gorilla-testing/app"
	"os"
)

func main() {
	port := os.Getenv("GORILLA_TESTING_PORT")

	gorillaTestingApp := app.NewGorillaTestingApp(port)

	if err := gorillaTestingApp.Run(); err != nil {
		return
	}
}
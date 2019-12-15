package app

import (
	"github.com/go-playground/log"
	"github.com/michaellihs/gorilla-testing/routing"
	"github.com/michaellihs/gorilla-testing/services"
	"net/http"
)

type GorillaTestingApp struct {
	port string
}

func NewGorillaTestingApp(port string) *GorillaTestingApp {
	return &GorillaTestingApp{port: port}
}

func (a *GorillaTestingApp) Run() error {
	log.Info("Starting GorillaTestingApp, listening on port ", a.port)
	helloService := services.NewService()
	router := routing.NewRouter(helloService)
	if err := http.ListenAndServe(":" + a.port, router); err != nil {
		log.Error(err)
	}
	return nil
}
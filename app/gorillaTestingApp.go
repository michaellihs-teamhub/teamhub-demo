package app

import (
	"github.com/michaellihs/gorilla-testing/services"
	"github.com/michaellihs/gorilla-testing/routing"
	"net/http"
)

type GorillaTestingApp struct {
	port string
}

func NewGorillaTestingApp(port string) *GorillaTestingApp {
	return &GorillaTestingApp{port: port}
}

func (a *GorillaTestingApp) Run() error {
	helloService := services.NewService()
	router := routing.NewRouter(helloService)
	if err := http.ListenAndServe(":" + a.port, router); err != nil {
		return err
	}
	return nil
}
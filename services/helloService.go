package services

import (
	"github.com/go-playground/log"
	"net/http"
)

type HelloService interface {
	Hello(http.ResponseWriter, *http.Request)
}

type helloService struct {

}

func NewService() HelloService {
	return &helloService{}
}

func (s *helloService) Hello(w http.ResponseWriter, r *http.Request) {
	log.Info("Processing request in helloService.Hello")
	if _, err := w.Write([]byte("hello")); err != nil {
		log.Error(err)
	}
}
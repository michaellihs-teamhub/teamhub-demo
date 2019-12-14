package services

import "net/http"

type HelloService interface {
	Hello(http.ResponseWriter, *http.Request)
}

type helloService struct {

}

func NewService() HelloService {
	return &helloService{}
}

func (s *helloService) Hello(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("hello")); err != nil {
		return
	}
}
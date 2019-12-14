package routing

import (
	"github.com/gorilla/mux"
	. "github.com/michaellihs/gorilla-testing/services"
	"net/http"
)

func NewRouter(helloService HelloService) http.Handler {
	router := mux.NewRouter()

	router.Path("/hello").Methods(http.MethodGet).HandlerFunc(helloService.Hello)

	return router
}
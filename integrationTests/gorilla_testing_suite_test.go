package integrationTests_test

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/michaellihs/gorilla-testing/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/michaellihs/gorilla-testing/routing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGorillaTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GorillaTesting Suite")
}

var _ = Describe("GIVEN the Gorilla Testing application up and running", func() {

	var handler http.Handler
	var server  *httptest.Server

	BeforeSuite(func() {
		helloService := services.NewService()
		handler = routing.NewRouter(helloService)
	})

	BeforeEach(func() {
		server = httptest.NewServer(handler)
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("WHEN it receives a GET request to '/hello'", func() {
		It("THEN it returns 'hello' and status code OK", func() {
			e := httpexpect.New(GinkgoT(), server.URL)

			e.GET("/hello").
				Expect().
				Status(http.StatusOK).
				Body().Equal("hello")
		})
	})

	Describe("WHEN it receives a GET request at an unknown endpoint", func() {
		It("THEN it returns the HTTP status code '404 NOT FOUND'", func() {
			e := httpexpect.New(GinkgoT(), server.URL)

			e.GET("/unknown-path").
				Expect().
				Status(http.StatusNotFound)
		})
	})
})

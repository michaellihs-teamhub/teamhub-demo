Gorilla Testing
===============

This repository contains a skeleton for a REST Service application written in Go. It provides

* routing with [Gorilla](https://www.gorillatoolkit.org/)
* BDD testing with [Ginkgo](http://onsi.github.io/ginkgo/)


Running the Tests
-----------------

In order to run the tests, run

    ginkgo -v -r


Running the Application
-----------------------

In order to run the application, run

    GORILLA_TESTING_PORT=8081 go run main.go

this will start a webserver listening on [http://localhost:8081/hello](http://localhost:8081/hello)


TODOs
-----

- [ ] Add consumer contract tests
- [ ] Figure out how to make the Ginkgo tests more `GIVEN WHEN THEN` style 

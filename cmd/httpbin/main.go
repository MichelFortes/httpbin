package main

import (
	"log"
	"michelfortes/httpbin/internal/constraints"
	"michelfortes/httpbin/internal/handlers"
	"net/http"
	"os"
)

func main() {
	port := getPort()
	http.Handle("/", &handlers.DefaultHandler{})
	http.Handle(constraints.EndpointProxy, &handlers.ProxyHandler{})
	http.Handle(constraints.EndpointHealth, &handlers.HealthHandler{})
	log.Printf(constraints.TextServerRunnning, port)
	log.Fatalln(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func getPort() string {
	port := os.Getenv(constraints.EnvServicePort)
	if len(port) != 0 {
		return port
	}
	return "8888"
}

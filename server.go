package main

import (
	"log"
	"net/http"
	"os"
	"softplan/httpbin-go/handlers"
)

func main() {
	port := getPort()
	http.Handle("/", &handlers.DefaultHandler{})
	http.Handle("/proxy", &handlers.ProxyHandler{})
	http.Handle("/__health", &handlers.HealthHandler{})
	log.Println("Server running on port", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

func getPort() string {
	port := os.Getenv("SERVER_PORT")
	if len(port) != 0 {
		return port
	}
	return "8888"
}

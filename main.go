package main

import (
	"fmt"
	"net/http"

	"github.com/kordano/johto/routers"
)


func main() {

	router := routers.InitRoutes()
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: router,
	}
	fmt.Println("Listening...")
	server.ListenAndServe()
}

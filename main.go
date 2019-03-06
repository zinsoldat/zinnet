package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	initlizeRoutes(router)

	port := ":3000"
	fmt.Println("\nListening on port " + port)
	http.ListenAndServe(port, router)
}

package main

import (
	"api/internal/config"
	"api/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	r := router.Generate()

	fmt.Printf("Listening on port %d", config.APIPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.APIPort), r))
}

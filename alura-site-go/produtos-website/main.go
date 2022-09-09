package main

import (
	"net/http"
	"github.com/alura-site-go/produtos-website/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}


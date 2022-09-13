package main

import (
	"github.com/alura-go/gin-api-rest/database"
	"github.com/alura-go/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}

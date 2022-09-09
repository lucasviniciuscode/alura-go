package main

import (
	"fmt"

	"github.com/alura-go/api-rest-go/database"
	"github.com/alura-go/api-rest-go/models"
	"github.com/alura-go/api-rest-go/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome", Historia: "Historia"},
		{Id: 2, Nome: "Nome", Historia: "Historia"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor rest em GO")
	routes.HandleRequest()
}

package main

import (
	"fmt"

	"github.com/api-rest-go/database"
	"github.com/api-rest-go/models"
	"github.com/api-rest-go/routes"
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

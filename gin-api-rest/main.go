package main

import (
	"github.com/gin-api-rest/database"
	"github.com/gin-api-rest/models"
	"github.com/gin-api-rest/routes"
)



func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "LuCAS" , CPF: "51549745184", RG: "54545152"},
		{Nome: "Ana" , CPF: "1234561", RG: "1234561"},
	}
	routes.HandleRequests()
}

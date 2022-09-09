package routes

import (
	"github.com/alura-go/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/:nome", controllers.Saudacao)
	r.Run()
}

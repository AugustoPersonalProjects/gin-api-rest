package routes

import (
	"github.com/AugustoPersonalProjects/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.Deleta)
	r.PATCH("/alunos/:id", controllers.Edita)
	r.Run()
}
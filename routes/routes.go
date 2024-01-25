package routes

import (
	"github.com/Monteiro712/go-gin-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandlerRequests()  {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeAlunos)
	r.POST("/alunos", controllers.CriarAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoID)
	r.DELETE("/alunos/:id", controllers.DeletarID)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarCPF)
	r.Run()
}
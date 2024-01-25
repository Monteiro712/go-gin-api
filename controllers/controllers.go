package controllers

import (
	"net/http"

	"github.com/Monteiro712/go-gin-api/db"
	"github.com/Monteiro712/go-gin-api/models"
	"github.com/gin-gonic/gin"
)

func ExibeAlunos(c *gin.Context)  {
	var alunos []models.Aluno
	db.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func CriarAlunos(c *gin.Context)  {
	var aluno models.Aluno
	//empacotar em json a struct aluno
	if err := c.ShouldBindJSON(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error()})
		return
	}
	if err := models.ValidarDadosAluno(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error()})
		return
	} 
	db.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoID(c *gin.Context)  {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	db.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"not found":"aluno não encontrado"})
			return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeletarID(c *gin.Context)  {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	db.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data":"aluno deletado"})
}

func EditarAluno(c *gin.Context)  {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	db.DB.First(&aluno, id)
	
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error()})
			return
	}
	if err := models.ValidarDadosAluno(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error()})
		return
	} 
	db.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscarCPF(c *gin.Context)  {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	db.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"not found":"aluno não encontrado"})
			return
	}
	c.JSON(http.StatusOK, aluno)
}
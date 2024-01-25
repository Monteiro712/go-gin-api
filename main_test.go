package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Monteiro712/go-gin-api/controllers"
	"github.com/Monteiro712/go-gin-api/db"
	"github.com/Monteiro712/go-gin-api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas;
}

func CriarAlunoMock()  {
	aluno := models.Aluno{Nome: "Jorgin teste", CPF: "90853114207"}
	db.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletarAlunoMock()  {
	var aluno models.Aluno
	db.DB.Delete(&aluno, ID)
}
func TestListarHandler(t *testing.T)  {
	db.ConectaComBancoDeDados()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	//fmt.Println(resposta.Body)
}

func TestRetornarCPFHandler(t *testing.T)  {
	db.ConectaComBancoDeDados()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscarCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/49874523012", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletarHandler(t *testing.T)  {
	db.ConectaComBancoDeDados()
	CriarAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletarID)
	pathBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaUmAlunoHandler(t *testing.T) {
	db.ConectaComBancoDeDados()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "47123456789"}
	valorJson, _ := json.Marshal(aluno)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "47123456789", alunoMockAtualizado.CPF)
	assert.Equal(t, "Nome do Aluno Teste", alunoMockAtualizado.Nome)
	fmt.Println(alunoMockAtualizado.CPF)
}

package controllers

import (
	"github.com/AgentTarik/api-go-gin/database"
	"github.com/AgentTarik/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Param("nome")
	c.JSON(200, gin.H{
		"API diz:": "Eai " + nome,
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"aluno": aluno,
	})
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")
	database.DB.First(&aluno, id)
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"aluno": aluno,
		"data":  "Aluno " + id + " deleted",
	})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID not found"})
		return
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&aluno).Updates(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Aluno not found",
		})
	}
	c.JSON(http.StatusOK, aluno)
}

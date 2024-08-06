package controllers

import (
	"github.com/AgentTarik/api-go-gin/database"
	"github.com/AgentTarik/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
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
	c.JSON(http.StatusOK, gin.H{
		"aluno": aluno,
	})
}

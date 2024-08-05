package controllers

import "github.com/gin-gonic/gin"

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Bruno",
	})
}

func Saudacao(c *gin.Context) {
	nome := c.Param("nome")
	c.JSON(200, gin.H{
		"API diz:": "Eai " + nome,
	})
}

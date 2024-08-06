package main

import (
	"github.com/AgentTarik/api-go-gin/database"
	"github.com/AgentTarik/api-go-gin/models"
	"github.com/AgentTarik/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Bruno", CPF: "22000000000", RG: "4700000000"},
		{Nome: "Jaum", CPF: "12300000000", RG: "5700000000"},
	}
	routes.HandleRequests()
}

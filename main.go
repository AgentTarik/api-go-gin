package main

import (
	"github.com/AgentTarik/api-go-gin/database"
	"github.com/AgentTarik/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}

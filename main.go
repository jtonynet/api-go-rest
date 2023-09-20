package main

import (
	"fmt"

	"github.com/jtonynet/api-go-rest/database"
	"github.com/jtonynet/api-go-rest/models"
	"github.com/jtonynet/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o server Rest com go")
	routes.HandleRequest()
}

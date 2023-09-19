package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jtonynet/api-go-rest/models"
	"github.com/jtonynet/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome1", Historia: "Historia 1"},
		{Nome: "Nome2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o server Rest com go")
	routes.HandleRequest()

	log.Fatal(http.ListenAndServe(":8000", nil))
}

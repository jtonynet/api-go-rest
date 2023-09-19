package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jtonynet/api-go-rest/models"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

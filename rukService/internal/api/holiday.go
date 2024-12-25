package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rukService/internal/models"
)

func GetPrazdnik(appId string) models.Holiday {
	var hol models.Holiday

	url := fmt.Sprintf("http://localhost:8082/get/holiday?app=%s", appId)
	r, err := http.Get(url)
	if err != nil {
		log.Println("Bad request", err)
		return models.Holiday{}
	}
	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(&hol); err != nil {
		log.Println(err)
		return models.Holiday{}
	}

	return hol
}

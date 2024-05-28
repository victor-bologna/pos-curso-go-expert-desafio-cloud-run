package web

import (
	"encoding/json"
	"net/http"

	"github.com/victor-bologna/pos-curso-go-expert-desafio-cloud-run/internal/usecase"
)

var WeatherService usecase.WeatherUsecaseInterface

func GetTempByCep(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if len(cep) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	response, err := WeatherService.Execute(cep)

	if err != nil {
		if err.Error() == "can not find zipcode" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

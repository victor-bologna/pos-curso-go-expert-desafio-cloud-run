package main

import (
	"github.com/victor-bologna/pos-curso-go-expert-desafio-cloud-run/internal/infra/web"
	"github.com/victor-bologna/pos-curso-go-expert-desafio-cloud-run/internal/infra/web/webserver"
	"github.com/victor-bologna/pos-curso-go-expert-desafio-cloud-run/internal/usecase"
)

func main() {
	webserver := webserver.NewWebServer("8080")
	webserver.AddHandler("GET", "/temperature", web.GetTempByCep)
	web.WeatherService = usecase.WeatherService{}
	webserver.Start()
}

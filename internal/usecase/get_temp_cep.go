package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/victor-bologna/pos-curso-go-expert-desafio-cloud-run/internal/infra/dto"
)

type WeatherUsecaseInterface interface {
	Execute(cep string) (OutputTempDTO, error)
}

type WeatherService struct{}

type OutputTempDTO struct {
	Temp_C float64 `json:"temp_C"`
	Temp_F float64 `json:"temp_F"`
	Temp_K float64 `json:"temp_K"`
}

func (ws WeatherService) Execute(cep string) (OutputTempDTO, error) {
	localidade, err := getLocalidade(cep)
	if err != nil {
		return OutputTempDTO{}, err
	}

	return getTemperatures(localidade)
}

func getLocalidade(cep string) (string, error) {
	response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	reader, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var viaCepResponse dto.ViaCepResponse
	err = json.Unmarshal(reader, &viaCepResponse)
	if err != nil {
		return "", err
	}
	if viaCepResponse.Erro {
		return "", errors.New("can not find zipcode")
	}

	return viaCepResponse.Localidade, nil
}

func getTemperatures(localidade string) (OutputTempDTO, error) {
	encodedCity := url.QueryEscape(localidade)

	apiKey := "06140e1756914c55b5a213915242605"
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, encodedCity)

	response, err := http.Get(url)
	if err != nil {
		return OutputTempDTO{}, err
	}
	defer response.Body.Close()

	reader, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputTempDTO{}, err
	}

	var weatherApiResponse dto.WeatherApiResponse
	err = json.Unmarshal(reader, &weatherApiResponse)
	if err != nil {
		return OutputTempDTO{}, err
	}

	return createOutputDTO(weatherApiResponse.Current.TempC)
}

func createOutputDTO(tempC float64) (OutputTempDTO, error) {
	return OutputTempDTO{
		Temp_C: tempC,
		Temp_F: (tempC*1.8 + 32),
		Temp_K: (tempC + 273),
	}, nil
}

package main

import (
	"net/http"
	"os"

	"github.com/felipemagrassi/weather-cep-api/internal/handler"
	"github.com/felipemagrassi/weather-cep-api/internal/service"
	"github.com/felipemagrassi/weather-cep-api/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	weatherApiKey, ok := os.LookupEnv("WEATHER_API_KEY")
	if !ok {
		weatherApiKey = ""
	}

	weather_service := service.NewWeatherApiService(weatherApiKey)
	cep_service := service.NewViaCepService()
	getTemperatureFromCepUseCase := usecase.NewGetTemperatureFromCepUseCase(cep_service, weather_service)
	getTemperatureHandler := handler.NewGetTemperatureHandler(getTemperatureFromCepUseCase)

	http.HandleFunc("/", getTemperatureHandler.Handle)
	http.ListenAndServe(":8080", nil)
}

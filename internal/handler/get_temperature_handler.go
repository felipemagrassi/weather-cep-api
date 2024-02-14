package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/felipemagrassi/weather-cep-api/internal/usecase"
)

type GetTemperatureHandler struct {
	getTemperatureFromCep *usecase.GetTemperatureFromCepUseCase
}

type GetTemperatureHandlerOutput struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func NewGetTemperatureHandler(getTemperatureFromCep *usecase.GetTemperatureFromCepUseCase) *GetTemperatureHandler {
	return &GetTemperatureHandler{getTemperatureFromCep: getTemperatureFromCep}
}

func (h *GetTemperatureHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cep, ok := h.getCep(r)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`invalid zipcode`))
		return
	}

	input := &usecase.GetTemperatureFromCepInput{Cep: cep}

	output, err := h.getTemperatureFromCep.Execute(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`can not find zipcode`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&GetTemperatureHandlerOutput{
		Celsius:    output.Celsius,
		Fahrenheit: output.Fahrenheit,
		Kelvin:     output.Kelvin,
	})

}

func (h *GetTemperatureHandler) getCep(r *http.Request) (string, bool) {
	cep := r.URL.Query().Get("cep")

	if cep == "" {
		return "", false
	}

	cepRegex := regexp.MustCompile(`^\d{5}-{0,1}\d{3}$`)
	if !cepRegex.Match([]byte(cep)) {
		return "", false
	}

	cep = strings.ReplaceAll(cep, "-", "")

	return cep, true
}

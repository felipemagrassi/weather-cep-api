package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type CepService interface {
	GetAddressByCep(ctx context.Context, cep string) (*ViaCepResponse, error)
}

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type ViaCepService struct {
	client *http.Client
}

var (
	CepServiceError = errors.New("error getting address")
)

func NewViaCepService() *ViaCepService {
	return &ViaCepService{client: &http.Client{}}
}

func (v *ViaCepService) GetAddressByCep(ctx context.Context, cep string) (*ViaCepResponse, error) {
	cep = strings.ReplaceAll(cep, "-", "")
	url := "https://viacep.com.br/ws/" + cep + "/json"

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := v.client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, CepServiceError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	viaCepResponse := &ViaCepResponse{}
	err = json.Unmarshal(body, &viaCepResponse)
	if err != nil {
		return nil, err
	}

	return viaCepResponse, nil
}

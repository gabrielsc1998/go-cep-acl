package brasil_api_gateway

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gabrielsc98/go-cep-acl/internal/domain"
)

type OutputDto struct {
	Code     string `json:"cep"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"neighborhood"`
	Address  string `json:"street"`
}

type Gateway struct {
}

func New() *Gateway {
	return &Gateway{}
}

func (v *Gateway) Get(cep string) (*domain.Address, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://brasilapi.com.br/api/cep/v1/"+cep,
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error")
	}

	outputDto := OutputDto{}
	err = json.NewDecoder(resp.Body).Decode(&outputDto)
	if err != nil {
		return nil, err
	}

	address := domain.NewAddress(
		outputDto.State,
		outputDto.City,
		outputDto.District,
		outputDto.Address,
	)
	return address, nil
}

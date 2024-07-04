package main

import (
	"fmt"

	"github.com/gabrielsc98/go-cep-acl/internal/application"
	brasil_api_gateway "github.com/gabrielsc98/go-cep-acl/internal/infra/brasilapi-gateway"
)

func main() {
	brasilAPIGateway := brasil_api_gateway.New()
	// viaCepGateway := viacep_gateway.New()
	addressACL := application.NewAddressACL(brasilAPIGateway)

	_, err := addressACL.GetAddressByCEP("14564")
	if err != nil {
		fmt.Println("1. CEP inválido:", err.Error() == "invalid cep")
	}

	resp, _ := addressACL.GetAddressByCEP("22010000")
	fmt.Println("2. Esperamos o endereço de Copacabana:", resp)

	isFromState, _ := addressACL.IsFromState("RJ", "22010000")
	fmt.Println("3. Esperamos 'true' para Copabacana ser de RJ:", isFromState)

	isFromState, _ = addressACL.IsFromState("SP", "22010000")
	fmt.Println("4. Esperamos 'false' para Copabacana ser de SP: ", isFromState)
}

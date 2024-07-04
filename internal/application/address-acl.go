package application

import (
	"errors"
	"regexp"
	"strings"

	"github.com/gabrielsc98/go-cep-acl/internal/domain"
)

type AddressACL struct {
	gateway domain.CEPGateway
}

func NewAddressACL(gateway domain.CEPGateway) *AddressACL {
	return &AddressACL{gateway: gateway}
}

func (acl *AddressACL) GetAddressByCEP(cep string) (*domain.Address, error) {
	address, err := acl.get(cep)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (acl *AddressACL) IsFromState(state string, cep string) (bool, error) {
	address, err := acl.get(cep)
	if err != nil {
		return false, err
	}
	return address.State == state, nil
}

func (acl *AddressACL) get(cep string) (*domain.Address, error) {
	if !isValidCEP(cep) {
		return nil, errors.New("invalid cep")
	}
	return acl.gateway.Get(normalizeCEP(cep))
}

func isValidCEP(cep string) bool {
	cepRegex := `^\d{5}-\d{3}$|^\d{8}$`
	match, _ := regexp.MatchString(cepRegex, cep)
	return match
}

func normalizeCEP(cep string) string {
	return strings.ReplaceAll(cep, "-", "")
}

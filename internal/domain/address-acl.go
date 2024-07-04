package domain

type AddressACL interface {
	GetAddressByCEP(cep string) (*Address, error)
	IsFromState(state string, cep string) (bool, error)
}

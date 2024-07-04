package domain

type CEPGateway interface {
	Get(cep string) (*Address, error)
}

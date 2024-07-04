package domain

type Address struct {
	State    string
	City     string
	District string
	Street   string
}

func NewAddress(state, city, district, street string) *Address {
	return &Address{
		State:    state,
		City:     city,
		District: district,
		Street:   street,
	}
}

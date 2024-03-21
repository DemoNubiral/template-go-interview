package source

import (
	"errors"
)

type HelloRepositoryInterface interface {
	GetName(parameter1 string, parameter2 string) (string, error)
}

type helloRepository struct {
}

func NewHelloRepository() helloRepository {
	return helloRepository{}
}

func (hr *helloRepository) GetName(parameter1 string, parameter2 string) (string, error) {

	if parameter1 == "123" {
		return "", errors.New("El parámetro 1 no es válido")
	}

	name := parameter1 + " " + parameter2

	return name, nil
}

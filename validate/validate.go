package validate

import "github.com/mkvlrn/docbr/internal"

func ValidateCPF(value string) (bool, error) {
	return internal.ValidateDocBR(internal.CPF, value)
}

func ValidateCNPJ(value string) (bool, error) {
	return internal.ValidateDocBR(internal.CNPJ, value)
}

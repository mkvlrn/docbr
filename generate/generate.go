package generate

import (
	"github.com/mkvlrn/docbr/internal"
)

func GenerateCPF(punctuation bool) (string, error) {
	return internal.GenerateDocBR(internal.CPF, punctuation)
}

func GenerateCNPJ(punctuation bool) (string, error) {
	return internal.GenerateDocBR(internal.CNPJ, punctuation)
}

package internal

import (
	"errors"
	"fmt"
)

const (
	errInvalidDocType = "invalid document type: must be CPF or CNPJ"
	errRegexExecution = "could not parse args for document: %w"
	errRegexMatch     = "document has invalid format and/or punctuation"
)

func newErrInvalidDocType() error {
	return errors.New(errInvalidDocType)
}

func newErrRegexExecution(err error) error {
	return fmt.Errorf(errRegexExecution, err)
}

func newErrRegexMatch() error {
	return errors.New(errRegexMatch)
}

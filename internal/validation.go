package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func validateFormat(doctype docType, value string) (bool, error) {
	regex := []string{`^\d{11}$`, `^\d{3}\.\d{3}\.\d{3}-\d{2}$`}
	if doctype == CNPJ {
		regex = []string{`^\d{14}$`, `^\d{2}\.\d{3}\.\d{3}\/\d{4}-\d{2}$`}
	}

	for _, r := range regex {
		valid, err := regexp.MatchString(r, value)
		if err != nil {
			return false, newErrRegexExecution(err)
		}

		if valid {
			return true, nil
		}
	}

	return false, nil
}

func removePunctuation(value string) string {
	clean := strings.ReplaceAll(value, "-", "")
	clean = strings.ReplaceAll(clean, "/", "")
	clean = strings.ReplaceAll(clean, ".", "")

	return clean
}

func validateVerifyingDigits(doctype docType, value string) (bool, error) {
	baseLength := 9
	if doctype == CNPJ {
		baseLength = 12
	}

	firstVd, err := generateVerifyingDigit(value[:baseLength])
	if err != nil {
		return false, err
	}

	secondVd, err := generateVerifyingDigit(value[:baseLength] + strconv.Itoa(firstVd))
	if err != nil {
		return false, err
	}

	return value[baseLength:] == fmt.Sprintf("%d%d", firstVd, secondVd), nil
}

func ValidateDocBR(doctype docType, value string) (bool, error) {
	if doctype != CPF && doctype != CNPJ {
		return false, newErrInvalidDocType()
	}

	validFormat, err := validateFormat(doctype, value)
	if err != nil {
		return false, err
	}

	if !validFormat {
		return false, newErrRegexMatch()
	}

	validDigits, err := validateVerifyingDigits(doctype, removePunctuation(value))
	if err != nil {
		return false, err
	}

	return validDigits, nil
}

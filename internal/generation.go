package internal

import (
	"math/rand"
	"strconv"
	"strings"
)

func generateRandomBase(doc docType) string {
	length := 9
	if doc == CNPJ {
		length = 12
	}

	base := []string{}
	for range length {
		base = append(base, strconv.Itoa(rand.Intn(10)))
	}

	return strings.Join(base, "")
}

func generateVerifyingDigit(base string) (int, error) {
	verifyingDigitsHash := map[int][]int{
		9:  {10, 9, 8, 7, 6, 5, 4, 3, 2},
		10: {11, 10, 9, 8, 7, 6, 5, 4, 3, 2},
		12: {5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2},
		13: {6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2},
	}
	mults := verifyingDigitsHash[len(base)]
	digits := strings.Split(base, "")

	sum := 0
	for i := range digits {
		digit, err := strconv.Atoi(digits[i])
		if err != nil {
			return 0, newErrRegexMatch()
		}
		sum += mults[i] * digit
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0, nil
	}

	return 11 - remainder, nil
}

func addPunctuation(doc docType, base string) string {
	if doc == CPF {
		return base[0:3] + "." + base[3:6] + "." + base[6:9] + "-" + base[9:]
	}

	return base[0:2] + "." + base[2:5] + "." + base[5:8] + "/" + base[8:12] + "-" + base[12:]
}

func GenerateDocBR(doc docType, punctuation bool) (string, error) {
	if doc != CPF && doc != CNPJ {
		return "", newErrInvalidDocType()
	}

	base := generateRandomBase(doc)
	firstVd, err := generateVerifyingDigit(base)
	if err != nil {
		return "", err
	}

	base += strconv.Itoa(firstVd)
	secondVd, err := generateVerifyingDigit(base)
	if err != nil {
		return "", err
	}

	base += strconv.Itoa(secondVd)
	if punctuation {
		return addPunctuation(doc, base), nil
	}

	return base, nil
}

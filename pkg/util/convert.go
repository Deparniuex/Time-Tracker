package util

import (
	"errors"
	"strconv"
	"strings"
)

func SplitPassport(passport string) (passportNumber int, passportSerie int, err error) {
	parts := strings.Split(passport, " ")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid format")
	}
	passportSerie, err1 := strconv.Atoi(parts[0])
	passportNumber, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("invalid numbers")
	}
	return passportNumber, passportSerie, nil
}

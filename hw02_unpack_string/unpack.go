package hw02unpackstring

import (
	"errors"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	runeString := []rune(str)
	result := ""
	lastChar := ""
	counterNumber := 0

	for i := 1; i <= len(runeString); i++ {
		nextSymbol := runeString[i-1 : i]
		nextSymbolIsNumber := checkingForANumber(nextSymbol)

		if nextSymbolIsNumber {
			counterNumber++
			if counterNumber > 1 || lastChar == "" {
				return "", ErrInvalidString
			}
			if string(nextSymbol) == "0" {
				result = result[:len(result)-1]
			}

			repeat, _ := strconv.Atoi(string(nextSymbol))
			for ii := 1; ii < repeat; ii++ {
				result += lastChar
			}
		} else {
			counterNumber = 0
			if nextSymbol[0] == '\n' {
				lastChar = "\\n"
				result += lastChar
			} else {
				lastChar = string(nextSymbol)
				result += lastChar
			}
		}
	}

	return result, nil
}

// Проверка на число
func checkingForANumber(symbol []rune) bool {
	if _, err := strconv.Atoi(string(symbol)); err == nil {
		return true
	} else {
		return false
	}
}

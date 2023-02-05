package hw02unpackstring

import (
	"errors"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// Place your code here.
	runeString := []rune(str)
	result := ""
	lastChar := ""
	counterNumber := 0
	var err error

	for i := 1; i <= len(runeString); i++ {
		nextSymbol := runeString[i-1 : i]

		//Проверка на число
		if _, err := strconv.Atoi(string(nextSymbol)); err == nil {
			counterNumber++
			if counterNumber > 1 {
				//panic(ErrInvalidString)
				err = ErrInvalidString
				break
			} else if lastChar == "" {
				//panic(ErrInvalidString)
				err = ErrInvalidString
				break
			} else if string(nextSymbol) == "0" {
				result = result[:len(result)-1]
				continue
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

	return result, err
}

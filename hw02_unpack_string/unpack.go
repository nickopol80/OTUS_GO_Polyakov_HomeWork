package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	lastChar := ""
	var result strings.Builder
	counterNumber := 0

	// Посимвольно прогоняем всё слово.
	for _, symbol := range str {
		// Условие если буква.
		if unicode.IsLetter(symbol) {
			counterNumber = 0
			lastChar = string(symbol)
			result.WriteString(string(lastChar))

			// Условие если не буква.
		} else if unicode.IsNumber(symbol) {
			counterNumber++
			numberRepeat, _ := strconv.Atoi(string(symbol))

			if counterNumber > 1 || lastChar == "" {
				return "", ErrInvalidString
			} else if numberRepeat == 0 {
				buf := result.String()
				buf = strings.TrimSuffix(buf, string(lastChar))
				result.Reset()
				result.WriteString(buf)
			} else {
				result.WriteString(strings.Repeat(string(lastChar), numberRepeat-1))
			}
		} else if unicode.IsSpace(symbol) {
			lastChar = strings.ReplaceAll(strconv.QuoteRuneToGraphic(symbol), "'", "")
			result.WriteString(lastChar)
		} else {
			counterNumber = 0
			lastChar = string(symbol)
			result.WriteString(string(lastChar))
		}

	}

	return result.String(), nil
}

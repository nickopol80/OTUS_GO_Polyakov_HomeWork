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
		switch {
		case unicode.IsLetter(symbol):
			counterNumber = 0
			lastChar = string(symbol)
			result.WriteString(lastChar)
		case unicode.IsNumber(symbol):
			counterNumber++
			numberRepeat, _ := strconv.Atoi(string(symbol))

			switch {
			case counterNumber > 1 || lastChar == "":
				return "", ErrInvalidString
			case numberRepeat == 0:
				buf := result.String()
				buf = strings.TrimSuffix(buf, lastChar)
				result.Reset()
				result.WriteString(buf)
			default:
				result.WriteString(strings.Repeat(lastChar, numberRepeat-1))
			}

		case unicode.IsSpace(symbol):
			lastChar = strings.ReplaceAll(strconv.QuoteRuneToGraphic(symbol), "'", "")
			result.WriteString(lastChar)
		default:
			counterNumber = 0
			lastChar = string(symbol)
			result.WriteString(lastChar)
		}
	}

	return result.String(), nil
}

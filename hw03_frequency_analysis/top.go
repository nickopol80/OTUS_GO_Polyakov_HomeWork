package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type stringEntity struct {
	Word        string
	Repetitions int
}

func Top10(str string) []string {
	// Place your code here.
	words := strings.Fields(str)
	counterWordsMap := make(map[string]int)

	// Подсчёт повторяющихся слов.
	for _, word := range words {
		_, matched := counterWordsMap[word]
		if matched {
			counterWordsMap[word]++
		} else {
			counterWordsMap[word] = 1
		}
	}

	// Перегоняю полученную статистику в слайс.
	var resultRating []stringEntity
	for str, num := range counterWordsMap {
		resultRating = append(resultRating, stringEntity{str, num})
	}

	// Сортировка в алфавитном порядке.
	sort.SliceStable(resultRating, func(i, j int) bool {
		return resultRating[i].Word < resultRating[j].Word
	})

	// Сортировка по убыванию количества повторов.
	sort.SliceStable(resultRating, func(i, j int) bool {
		return resultRating[i].Repetitions > resultRating[j].Repetitions
	})

	// Формирование ответа.
	result := []string{}
	counter := 10
	for _, entity := range resultRating {
		if counter > 0 {
			result = append(result, entity.Word)
			counter--
		} else {
			return result
		}
	}

	return result
}

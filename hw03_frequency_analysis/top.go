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
		counterWordsMap[word]++
	}

	// Создаю новый слайс с заданнной длиной.
	resultRating := make([]stringEntity, len(counterWordsMap))

	// Перегоняю полученную статистику в слайс.
	for str, num := range counterWordsMap {
		resultRating = append(resultRating[1:], stringEntity{str, num})
	}

	// Сортировка.
	sort.Slice(resultRating, func(i, j int) bool {
		if resultRating[i].Repetitions != resultRating[j].Repetitions {
			return resultRating[i].Repetitions > resultRating[j].Repetitions
		}

		return resultRating[i].Word < resultRating[j].Word
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

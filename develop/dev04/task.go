package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"sort"
	"strconv"
	"strings"
)

func anagramma(a *[]string) *map[string][]string {
	var arStringIntAll []string
	map1 := make(map[string]int)
	map2 := make(map[string][]string)
	map3 := make(map[string][]string)
	for _, v := range *a {
		word := strings.ToLower(v)
		runes := []rune(word)
		var arRuneInt []int
		for _, v1 := range runes {
			arRuneInt = append(arRuneInt, int(v1))
		}
		sort.Ints(arRuneInt)
		var arStringIntPart []string
		for _, v2 := range arRuneInt {
			arStringIntPart = append(arStringIntPart, strconv.Itoa(v2)) //string
		}
		myStr := strings.Join(arStringIntPart, "")
		arStringIntAll = append(arStringIntAll, myStr)

	}
	for _, v3 := range arStringIntAll {
		map1[v3]++
	}
	for i4, v4 := range map1 {
		if v4 > 1 {
			for i5, v5 := range arStringIntAll {
				if i4 == v5 {
					map2[i4] = append(map2[i4], strings.ToLower((*a)[i5]))
				}
			}
		}
	}
	for _, t4 := range map2 {
		map3[t4[0]] = t4
	}
	for _, t5 := range map3 {
		sort.Strings(t5)
	}
	return &map3
}


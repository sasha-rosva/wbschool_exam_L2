package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"fmt"
	"strconv"

	"strings"
)

func decode(s string) (string, error) {
	sb := new(strings.Builder)
	runes := []rune(s)
	i := 0
	for _, v := range runes {
		if v < 48 || v > 57 {
			i++
			break
		}
	}
	if i == 0 && s != "" {
		return "", errors.New("некорректная строка")
	}
	var a, aa int32
	map1 := make(map[int32]int)
	for _, v := range runes {
		if a == 92 && map1[92] == 1 {
			sb.WriteString(string(v))
			a = v
			map1[92] = 0
			continue
		}
		if v >= 48 && v <= 57 && a != 0 && (map1[92] == 2 || map1[92] == 0) {
			if v == 48 && (a < 48 || a > 57) {
				return "", errors.New("число начинается с нуля")
			}
			if a > 48 && a < 58 && aa != 92 {
				num, _ := strconv.Atoi(string(a) + string(v))
				num1, _ := strconv.Atoi(string(a))
				for i := 0; i < num-num1; i++ {
					sb.WriteString(string(aa))
				}
			} else {
				num, _ := strconv.Atoi(string(v))
				for i := 0; i < num-1; i++ {
					sb.WriteString(string(a))
				}
			}
		} else {
			if v != 32 && v != 92 {
				a = v
				sb.WriteString(string(a))
			}
		}
		if v == 92 {
			map1[92]++
		}
		a = v
		if v < 48 || v > 57 {
			aa = v
		}
	}
	return sb.String(), nil
}

func main() {
	if y, err := decode(`a12\3\4\5\\6\5`); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(y)
	}
}

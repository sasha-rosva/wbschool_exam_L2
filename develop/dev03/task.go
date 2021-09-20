package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortT(s []string) []string {
	sort.Strings(s)
	return s

}

// -r — сортировать в обратном порядке
func sortR(s []string) []string {
	sort.Strings(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// -u — не выводить повторяющиеся строки
func sortU(s []string) []string {
	map1 := make(map[string]int)
	var slice []string
	for _, v := range s {
		map1[v]++
		if map1[v] == 1 {
			slice = append(slice, v)
		}
	}
	return sortT(slice)
}

// -k — указание колонки для сортировки
func sortK(s []string, n int) ([]string, error) {
	map1 := make(map[interface{}]string)
	var arr, arr1 []string
	for _, v := range s {
		a := strings.Split(v, " ")
		if n > len(a) || n < 1 {
			return nil, fmt.Errorf("index out of range [%d] with length %d", n, len(a))

		}
		map1[a[n-1]] = v
		arr = append(arr, a[n-1])
	}

	sort.Strings(arr)
	for _, v := range arr {
		arr1 = append(arr1, map1[v])
	}
	return arr1, nil
}

// -n — сортировать по числовому значению
func sortN(s []string, n int) ([]string, error) {
	map1 := make(map[string][]string)
	var arr []int
	var arr1, arr2 []string
	for _, v := range s {
		a := strings.Split(v, " ")
		if n > len(a) || n < 1 {
			return nil, fmt.Errorf("index out of range [%d] with length %d", n, len(a))

		}
		if _, err := strconv.Atoi(a[n-1]); err != nil {
			return nil, fmt.Errorf("the value of column [%d] is not int", n)

		}
		map1[a[n-1]] = append(map1[a[n-1]], v)
		i, _ := strconv.Atoi(a[n-1])
		arr = append(arr, i)
	}
	sort.Ints(arr)
	// числа превращаем в числа-строки
	for _, vv := range arr {
		arr1 = append(arr1, strconv.Itoa(vv))
	}
	for _, v := range arr1 {
		if vv, ok := map1[v]; ok {
			arr2 = append(arr2, vv[0])
			if len(map1[v]) > 1 {
				map1[v] = map1[v][1:]
			}
		}
	}
	return arr2, nil
}

// -c — проверять отсортированы ли данные
func sortC(s []string) bool {
	tmp := make([]string, len(s))
	copy(tmp, s)
	sort.Strings(s)
	for i, v := range s {
		if v != tmp[i] {
			return false
		}
	}
	return true

}

// // -b — игнорировать хвостовые пробелы
func sortB(s []string) []string {
	sb := new(strings.Builder)
	var f []string
	for _, v := range s {
		y := strings.Fields(v)
		for i, vv := range y {
			sb.WriteString(vv)
			if i != len(y)-1 {
				sb.WriteString(" ")
			}
		}
		f = append(f, sb.String())
		sb.Reset()
	}

	return sortT(f)
}

// -h — сортировать по числовому значению с учётом суффиксов
func sortH(s []string, n int) ([]string, error) {
	var arr11 []string
	map11 := make(map[interface{}]string)
	for _, v := range s {
		a := strings.Split(v, " ")
		map11[a[n-1]] = v
		i := a[n-1]
		arr11 = append(arr11, i)
	}
	map1 := make(map[interface{}][]string)
	var arr []int
	for _, v := range s {
		t := strings.Split(v, " ")
		runes := []rune(t[n-1])
		if runes[0] < 48 || runes[0] > 57 {
			return nil, errors.New("prefix is not int")
		}
		for ii, vv := range runes {
			if vv < 28 || vv > 57 {
				num, _ := strconv.Atoi(string(runes[:ii]))
				arr = append(arr, num)
				map1[num] = append(map1[num], string(runes[ii:]))
				break
			}
		}
	}
	sort.Ints(arr)
	sb := new(strings.Builder)
	var arr1, arr2 []string
	var ss int
	for _, v := range arr {
		if v != ss {
			arr1 = map1[v]
			sort.Strings(arr1)
			for _, vv := range arr1 {
				sb.WriteString(strconv.Itoa(v))
				sb.WriteString(vv)
				arr2 = append(arr2, map11[sb.String()])
				sb.Reset()
			}
			ss = v
		}
	}
	return arr2, nil
}

// -M — сортировать по названию месяца
func sortM(s []string, n int) ([]string, error) {
	var arr11 []string
	sb := new(strings.Builder)
	map1 := make(map[string]int)
	map2 := make(map[int]string)
	map11 := make(map[interface{}][]string)
	var ar1 []int
	var ar2 []string
	ar := []string{"январь", "февраль", "март", "апрель", "май", "июнь", "июль", "август",
		"сентябрь", "октябрь", "ноябрь", "декабрь"}
	for i, v := range ar {
		map1[v] = i
	}

	for ii, v := range s {
		a := strings.Split(v, " ")
		if ii == 0 {
			w := 0
			for _, vv := range ar {
				if p := strings.ToLower(a[n-1]); p != vv {
					w++
				}
			}
			if w == 12 {
				return nil, fmt.Errorf("the value of column [%d] is not month", n)
			}
		}
		map11[a[n-1]] = append(map11[a[n-1]], v)
		i := a[n-1]
		arr11 = append(arr11, i)
	}

	for _, v := range arr11 {
		vv := map1[strings.ToLower(v)]
		map2[vv] = v
		ar1 = append(ar1, map1[strings.ToLower(v)])
	}
	sort.Ints(ar1)
	for _, v := range ar1 {
		strArr := map11[map2[v]]
		for i, vv := range strArr {
			sb.WriteString(vv)
			if i != len(strArr)-1 {
				sb.WriteString(" ")
			}
		}
		ar2 = append(ar2, sb.String())
		sb.Reset()
	}

	return ar2, nil
}
func openFile(s string) ([]string, error) {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}
	aa := strings.ReplaceAll(string(data), "\r", "")
	a := strings.Split(aa, "\n")
	for i, v := range a {
		runes := []rune(v)
		if runes[0] == 65279 {
			a[i] = string(runes[1:])
		}
	}
	return a, nil
}
func open(s string) []string {
	aa := strings.ReplaceAll(s, "\r", "")
	a := strings.Split(aa, "\n")
	return a
}
func initParse() (*int, []*bool) {
	var arr []*bool
	k := flag.Int("k", 0, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	m := flag.Bool("m", false, "сортировать по названию месяца (Январь, Февраль...)")
	b := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	c := flag.Bool("c", false, "проверять отсортированы ли данные")
	h := flag.Bool("h", false, "сортировать по числовому значению с учётом суффиксов")
	arr = append(arr, n, r, u, m, b, c, h)
	return k, arr
}

type caseSort struct {
	k *int
	b []*bool
}

func main() {
	sss := new(caseSort)
	sss.k, sss.b = initParse()
	flag.Parse()
	var text string
	var array []string
	var err error
	if len(flag.Args()) != 0 {
		text = flag.Args()[0]
		array, err = openFile(text)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Вы не указали имя файла! Ниже у Вас есть возможность ввести текст вручную!")
		fmt.Print("Введите текст: ")
		_, _ = fmt.Scan(&text)
		array = open(text)
	}

	if *sss.b[5] {
		ok := sortC(array)
		if ok {
			fmt.Println("отсортирован")
			os.Exit(0)
		} else {
			fmt.Println("не отсортирован")
			os.Exit(0)
		}
	}
	array = sortT(array)
	if *sss.b[4] {
		array = sortB(array)
	}
	if *sss.b[2] {
		array = sortU(array)
	}
	if *sss.k != 0 {
		array, err = sortK(array, *sss.k)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
	a := []int{0, 3, 6}
	aa := 0
	for _, v := range a {
		if *sss.b[v] {
			aa++
		}
		if aa > 1 {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", errors.New("из ключей (h,m,n) можно выбрать лишь один"))
			os.Exit(1)
		}
	}
	if *sss.b[0] {
		if *sss.k == 0 {
			fmt.Println("Забыли указать номер столбца!")
			fmt.Print("Номер столбца: ")
			_, err := fmt.Scan(sss.k)
			if err != nil {
				fmt.Printf("error: %v", err)
				os.Exit(1)
			}
		}
		array, err = sortN(array, *sss.k)
		if err != nil {
			fmt.Printf("error: %v", err)
			os.Exit(1)
		}
	}
	if *sss.b[3] {
		if *sss.k == 0 {
			fmt.Println("Забыли указать номер столбца!")
			fmt.Print("Номер столбца: ")
			_, err := fmt.Scan(sss.k)
			if err != nil {
				fmt.Printf("error: %v", err)
				os.Exit(1)
			}
		}
		array, err = sortM(array, *sss.k)
		if err != nil {
			fmt.Printf("error: %v", err)
			os.Exit(1)
		}
	}
	if *sss.b[6] {
		if *sss.k == 0 {
			fmt.Println("Забыли указать номер столбца!")
			fmt.Print("Номер столбца: ")
			_, err := fmt.Scan(sss.k)
			if err != nil {
				fmt.Printf("error: %v", err)
			}
		}
		array, err = sortH(array, *sss.k)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
	}
	if *sss.b[1] {
		array = sortR(array)
	}

	for _, v := range array {
		fmt.Println(v, len(v))
	}

}

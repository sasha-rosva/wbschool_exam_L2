package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	f    = flag.Int("f", 0, `"fields" - выбрать поля (колонки)`)
	d    = flag.String("d", `\t`, `"delimiter" - использовать другой разделитель`)
	s    = flag.Bool("s", false, `"separated" - только строки с разделителем`)
	file string
	err  error
	a    [][]string
)

func openFile(file string, sss *cut) ([][]string, error) {
	var ttt [][]string
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	aa := strings.ReplaceAll(string(data), "\r", "")
	arrayStR := strings.Split(aa, `\n`)
	for i, v := range arrayStR {
		if v != "" {
			runes := []rune(v)
			if runes[0] == 65279 {
				arrayStR[i] = string(runes[1:])
			}
		}
	}
	for _, vv := range arrayStR {
		if *sss.s {
			if ok := strings.Contains(vv, *sss.d); ok {
				arrayStR2 := strings.Split(vv, *sss.d)
				ttt = append(ttt, arrayStR2)
			}
		} else {
			arrayStR2 := strings.Split(vv, *sss.d)
			ttt = append(ttt, arrayStR2)
		}
	}

	return ttt, nil
}
func open(s string, sss *cut) [][]string {
	var ttt [][]string
	aa := strings.ReplaceAll(s, "\r", "")
	arrayStR := strings.Split(aa, `\n`)
	for i, v := range arrayStR {
		if v != "" {
			runes := []rune(v)
			if runes[0] == 65279 {
				arrayStR[i] = string(runes[1:])
			}
		}
	}
	for _, vv := range arrayStR {
		if *sss.s {
			if ok := strings.Contains(vv, *sss.d); ok {
				arrayStR2 := strings.Split(vv, *sss.d)
				ttt = append(ttt, arrayStR2)
			}
		} else {
			arrayStR2 := strings.Split(vv, *sss.d)
			ttt = append(ttt, arrayStR2)
		}
	}
	return ttt
}

type cut struct {
	f *int
	d *string
	s *bool
}

func main() {
	sss := new(cut)
	sss.f = f
	sss.d = d
	sss.s = s
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		file = args[0]
		a, err = openFile(file, sss)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Вы не указали ключевое слово или имя файла! Ниже у Вас есть возможность ввести текст вручную!")
		fmt.Print("Введите текст: ")
		_, _ = fmt.Scan(&file)
		a = open(file, sss)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
	for _, v := range a {
		if len(v) < *sss.f {
			_, _ = fmt.Fprintf(os.Stderr, "error: index out of range [%d] with length %d\n", *sss.f, len(v))
			os.Exit(1)
		}
	}
	fmt.Println(a)

	for _, v := range a {
		for ii, vv := range v {
			if ii == *sss.f-1 {
				fmt.Println(vv)
			}
		}
	}
}

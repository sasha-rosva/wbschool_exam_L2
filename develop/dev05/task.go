package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	a     = flag.Int("A", 0, `"after" печатать +N строк после совпаденияи`)
	b     = flag.Int("B", 0, `"before" печатать +N строк до совпадения`)
	cc    = flag.Int("C", 0, `"context" (A+B) печатать ±N строк вокруг совпадения`)
	c     = flag.Bool("c", false, `"count" (количество строк)`)
	i     = flag.Bool("i", false, `"ignore-case" (игнорировать регистр)`)
	v     = flag.Bool("v", false, `"invert" (вместо совпадения, исключать)`)
	f     = flag.Bool("F", false, `"fixed", точное совпадение со строкой, не паттерн`)
	n     = flag.Bool("n", false, `"line num", печатать номер строки`)
	file  string
	text  string
	index int
)

func openFile(file string) ([]string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	aa := strings.ReplaceAll(string(data), "\r", "")
	arrayStR := strings.Split(aa, "\n")
	for i, v := range arrayStR {
		if v != "" {
			runes := []rune(v)
			if runes[0] == 65279 {
				arrayStR[i] = string(runes[1:])
			}
		}
	}
	return arrayStR, nil
}
func grepA(arr []string, i int, n int, b bool) {
	a := make(map[int]int)
	for t := 0; t <= n && i+t < len(arr)-1; t++ {
		if vv := a[i+t]; vv == 0 {
			if b {
				fmt.Println(arr[i+t], "номер строки:", i+t)
			} else {
				fmt.Println(arr[i+t])
			}
			a[i+t]++
		}
	}

}

func grepB(arr []string, i int, n int, b bool) {
	a := make(map[int]int)
	for t := n; t >= 0; t-- {
		if vv := a[i-t]; vv == 0 && (i-t) >= 0 {
			if b {
				fmt.Println(arr[i-t], "номер строки:", i-t)
			} else {
				fmt.Println(arr[i-t])
			}
			a[i-t]++
		}
	}
}

func grepCC(arr []string, i int, n int, b bool) {
	a := make(map[int]int)

	for t := n; t >= 0; t-- {
		if vv := a[i-t]; vv == 0 && (i-t) >= 0 {
			if b {
				fmt.Println(arr[i-t], "номер строки:", i-t)
			} else {
				fmt.Println(arr[i-t])
			}
			a[i-t]++
		}
	}
	for t := 0; t <= n && i+t < len(arr)-1; t++ {
		if vv := a[i+t]; vv == 0 {
			if b {
				fmt.Println(arr[i+t], "номер строки:", i+t)
			} else {
				fmt.Println(arr[i+t])
			}
			a[i+t]++
		}
	}

}

func grepC(arr []string, text string) {
	t := 0
	for _, v := range arr {
		contain := strings.Contains(v, text)
		if contain {
			t++
		}
	}
	fmt.Println(t)
}

type grep struct {
	flagsB []*bool
	flagsN []*int
}

func newt(arr []string, text string, grep *grep) []int {
	var as []int
	var arrLow []string
	var contain bool
	if *grep.flagsB[1] {
		text = strings.ToLower(text)
		for _, vv := range arr {
			vv = strings.ToLower(vv)
			arrLow = append(arrLow, vv)
		}
		arr = arrLow
	}
	for i, v := range arr {
		if v != "" {
			if !*grep.flagsB[3] {
				contain = strings.Contains(v, text)
			} else {
				contain = v == text
			}
			if !*grep.flagsB[2] {
				if contain {
					as = append(as, i)
				}
			} else {
				if !contain {
					as = append(as, i)
				}
			}
		}
	}
	return as
}
func checkFlags(grep *grep) (int, error) {
	a := []int{0, 1, 2}
	var aa, aaa int
	for i, v := range a {
		if *grep.flagsN[v] != 0 {
			aa++
			aaa = i + 1
		}
		if aa > 1 {
			return 0, errors.New("из ключей (A,B,C) можно выбрать лишь один")
		}
	}
	return aaa, nil
}
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		text = args[0]
		file = args[1]
	} else {
		fmt.Println("Вы не указали ключевое слово или имя файла!")
		os.Exit(1)
	}
	sss := new(grep)
	sss.flagsN = append([]*int{}, a, b, cc)
	sss.flagsB = append([]*bool{}, c, i, v, f, n)
	arrayStR, err := openFile(file)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if *sss.flagsB[0] {
		grepC(arrayStR, text)
		os.Exit(0)
	}

	index, err = checkFlags(sss)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	nnn := newt(arrayStR, text, sss)
	for _, p := range nnn {
		if index != 0 {
			switch index - 1 {
			case 0:
				grepA(arrayStR, p, *sss.flagsN[index-1], *sss.flagsB[4])
			case 1:
				grepB(arrayStR, p, *sss.flagsN[index-1], *sss.flagsB[4])
			case 2:
				grepCC(arrayStR, p, *sss.flagsN[index-1], *sss.flagsB[4])
			}
		} else {
			if *sss.flagsB[4] {
				fmt.Println(arrayStR[p], "номер строки:", p)
			} else {
				fmt.Println(arrayStR[p])
			}
		}
	}

}

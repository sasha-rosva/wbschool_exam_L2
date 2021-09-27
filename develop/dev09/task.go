package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	output = flag.String("o", "", "имя файла, в который будут сохранены полученные данные")
)

func wget(url, fileName string) {
	resp := getResponse(url)
	writeToFile(fileName, resp)
}

func getResponse(url string) *http.Response {
	tr := new(http.Transport)
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	errorChecker(err)
	return resp
}

func writeToFile(fileName string, resp *http.Response) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	errorChecker(err)
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	errorChecker(err)
}

func errorChecker(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
func main() {
	flag.Parse()
	args := flag.Args()
	var s string
	if len(args) < 1 {
		fmt.Println("Введите адрес ссылки:")
		_, err := fmt.Scan(&s)
		errorChecker(err)
	} else {
		s = args[0]
	}
	if *output == "" {
		as := strings.Split(s, "/")
		if len(as[len(as)-1]) > 3 {
			*output = as[len(as)-1]
		} else {
			fmt.Println("Выберите имя файла и его расширение:")
			_, err := fmt.Scan(output)
			errorChecker(err)
		}
	}
	wget(s, *output)
}


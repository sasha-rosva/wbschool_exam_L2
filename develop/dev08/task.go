package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"

	"strings"
)

func startClient(addr string) error {
	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", addr)
	if err != nil {
		conn, err = net.Dial("udp", addr)
		if err != nil {
			fmt.Printf("Can't connect to server: %s\n", err)
			return err
		}
	}
	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		fmt.Printf("Connection error: %s\n", err)
		return err
	}
	return nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		dir, _ := os.Getwd()
		fmt.Printf("%s$ > ", dir)
		input, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		if err = execInput(input); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}

var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	var cmd *exec.Cmd

	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":

		if len(args) < 2 {
			return ErrNoPath
		}

		return os.Chdir(args[1])
	case "netcat":
		return startClient(fmt.Sprintf("%s:%s", args[1], args[2]))

	case "exit":
		os.Exit(0)
	}

	if !strings.Contains(input, "|") {
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		cmd = exec.Command("bash", "-c", input)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

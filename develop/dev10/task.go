package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	timeout time.Duration
	host    string
	port    string
)

func init() {
	flag.Usage = func() {
		fmt.Println("Usage flags: [--timeout] host port")
		flag.PrintDefaults()
	}

	flag.DurationVar(&timeout, "timeout", time.Second*10, "timeout")

	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	host = args[0]
	port = args[1]
}

func read(conn net.Conn, cancelFunc context.CancelFunc) {
	scanner := bufio.NewScanner(conn)
	for {
		if !scanner.Scan() {
			log.Printf("соединение было прервано")
			cancelFunc()
			return
		}
		text := scanner.Text()
		fmt.Printf("%s\n", text)
	}
}

func write(conn net.Conn, cancelFunc context.CancelFunc) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			log.Println("CANNOT STDIN SCAN")
			cancelFunc()
			return
		}
		str := scanner.Text()

		sl := strings.Split(fmt.Sprintf("% x", str), " ")
		for _, u := range sl {
			if u == "04" {
				log.Println("вы нажали ctrl+D. telnet клиент будет закрыт!")
				cancelFunc()
				return
			}

		}
		_, err := conn.Write([]byte(fmt.Sprintln(str)))
		if err != nil {
			log.Println("ошибка при отправке на сервер", err)
			cancelFunc()
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	go func() {
		sgn := <-signalChan
		log.Println(sgn.String())
		cancel()
	}()

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	go read(conn, cancel)
	go write(conn, cancel)

	<-ctx.Done()
	log.Println("telnet клиент закрыт")
}

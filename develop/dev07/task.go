package main

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

import (
	"fmt"
	"sync"
	"time"
)

func manyToOne(channels ...<-chan interface{}) <-chan interface{} {
	var group sync.WaitGroup
	output := make(chan interface{}, 1)
	group.Add(len(channels))
	for i := range channels {
		go func(input <-chan interface{}) {
			for val := range input {
				output <- val
			}
			group.Done()
		}(channels[i])
	}
	go func() {
		group.Wait()
		close(output)
	}()
	return output
}
func main() {
	sig := func(second int, n int) <-chan interface{} {
		c := make(chan interface{}, 1)
		go func() {
			defer close(c)
			for i := 0; i < second; i++ {
				c <- n
				time.Sleep(time.Second)
			}

		}()
		return c
	}
	start := time.Now()
	out := manyToOne(
		sig(20, 0),
		sig(15, 1),
		sig(10, 2),
		sig(5, 3),
		sig(1, 4),
	)
	for vvv := range out {
		fmt.Println(vvv)
	}
	fmt.Printf("fone after %v", time.Since(start))
}

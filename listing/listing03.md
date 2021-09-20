Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```go
package main

import (
	"fmt"
	"os"
)
//Интерфейс по своей сути - это всего лишь набор данных о типах
//(интерфейса и типа переменной внутри него) и указатель на, собственно, саму переменную
//со статическим (конкретным) типом (поле data в iface). Таким образом в err содержится информация о
// типе переменной. Тип переменной - это  *os.PathError. Соответственно err уже не будет nil.
func Foo() error {
	var err *os.PathError = nil
	return err
}
func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}

```

package pattern

/*
Паттерн фабричный метод - основа всех порождающих паттернов. В Golang можно реализовать лишь базовую версию паттерна - простая фабрика.
Плюсы: выделяет код производства продуктов в одно место, упрощая поддержку кода. 
*/
import (
	"errors"
	"fmt"
)

type Player interface {
	GetName() string
	GetAge() int
	GetGoals() int
}

func getFootballPlayer(role string) (Player, error) {
	if role == "goalkeeper" {
		return newGoalkeeper(), nil
	}
	if role == "field player" {
		return newFieldPlayer(), nil
	}
	return nil, errors.New("wrong player type passed")
}
func main() {
	player1, _ := getFootballPlayer("goalkeeper")
	player2, _ := getFootballPlayer("field player")
	printDetails(player1)
	fmt.Println()
	printDetails(player2)

	fmt.Println()
	player4, err := getFootballPlayer("coach")
	if err == nil {
		printDetails(player4)
	} else {
		fmt.Println(err)
	}

}

func printDetails(p Player) {
	fmt.Printf("Name: %s\n", p.GetName())
	fmt.Printf("Age: %d\n", p.GetAge())
	fmt.Printf("Goals: %d\n", p.GetGoals())
}

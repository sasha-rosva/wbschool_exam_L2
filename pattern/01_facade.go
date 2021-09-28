package pattern
/* Суть паттерна фасад - это скрыть сложные детали реализации, чтобы конечному пользователю было легко им пользоваться.
Паттерн используется, когда нужно предоставить простой интерфейс к сложной подсистеме.
Пример: "Умный дом"
*/

import (
	"fmt"
	"time"
)

type SmartHouseFacade struct {
	Water Water
	Light Light
	Tv    Tv
}
type Water struct {}
type Light struct {}
type Tv struct {}
func NewSmartHouseFacade()*SmartHouseFacade{
	return &SmartHouseFacade{}
}
func (w SmartHouseFacade) IAmHome(){
	fmt.Println("Запущен режим: 'Я дома!'")
	w.Water.TurnOnWater()
	w.Light.SwitchOnLight()
	w.Tv.ActivateTv()
}
func (w SmartHouseFacade) IGoAway(){
	fmt.Println("Запущен режим: 'Я ушел!'")
	w.Water.TurnOffWater()
	w.Light.SwitchOffLight()
	w.Tv.DeactivateTv()
}
func (w Water) TurnOnWater(){
	fmt.Println("Вода включена!")
}
func (w Water) TurnOffWater(){
	fmt.Println("Вода выключена!")
}
func (w Light) SwitchOnLight(){
	fmt.Println("Свет включен!")
}
func (w Light) SwitchOffLight(){
	fmt.Println("Свет выключен!")
}
func (w Tv) ActivateTv(){
	fmt.Println("Телевизор включен!")
}
func (w Tv) DeactivateTv(){
	fmt.Println("Телевизор выключен!")
}
func main(){
	MySmartHouseFacade:=NewSmartHouseFacade()
	MySmartHouseFacade.IAmHome()
	time.Sleep(3*time.Second)
	MySmartHouseFacade.IGoAway()
}

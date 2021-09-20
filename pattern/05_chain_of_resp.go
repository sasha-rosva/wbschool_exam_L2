package pattern

/*
Паттерн цепочка обязанностей позволяет передавать запрос по цепочке потенциальных обработчиков, пока один из них не обработает запрос. 
Плюсы: избавляет от жёсткой привязки отправителя запроса к его получателю. Минусы: запрос теряется, если не удовлетворяет ни одному обработчику.
*/
import "fmt"

type Case interface {
	execute(*Recipe)
	setNext(Case)
}
type Recipe struct {
	WantRecipe string
}

type Sister struct {
	next Case
	myRecipe string
}
func (s *Sister) execute(r *Recipe){
	if r.WantRecipe==s.myRecipe {fmt.Println("I am sister. Yes, i have this recipe!")} else{s.next.execute(r)}
}
func (s *Sister) setNext(next Case) {
	s.next = next
}
type Internet struct{
	next Case
	myRecipe string
}
func (i *Internet) execute(r *Recipe){
	if r.WantRecipe==i.myRecipe {fmt.Println("I am internet. Yes, i have this recipe!")} else{fmt.Println("Nobody knows this recipe!:(")}
}
func (i *Internet) setNext(next Case) {
	i.next = next
}

type Friend struct{
	next Case
	myRecipe string
}
func (f *Friend) execute(r *Recipe){
	if r.WantRecipe==f.myRecipe {fmt.Println("I am friend. Yes, i have this recipe!")} else{f.next.execute(r)}
}
func (f *Friend) setNext(next Case) {
	f.next = next
}

func main(){
	internet:=&Internet{myRecipe: "Оливье"}

	sister:=&Sister{myRecipe: "Пицца"}
	sister.setNext(internet)

	friend:=&Friend{myRecipe: "Яичница"}
	friend.setNext(sister)

	recipe:=&Recipe{WantRecipe: "Оливье"}
	friend.execute(recipe)

}

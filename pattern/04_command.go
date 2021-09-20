package pattern

/*
Паттерн команда позволяет заворачивать запросы или простые операции в отдельные объекты. Эти запросы можно поставить в очередь, отменить, логировать.
Плюсы: убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют; позволяет реализовать простую отмену и повтор 
операций; позволяет реализовать отложенный запуск операций; позволяет собирать сложные команды из простых. Минусы: усложняет код программы из-за введения множества дополнительных структур.
*/
import (
	"errors"
	"fmt"
	"sync"
)
type Database struct {
	repo map[string]int
	mutex sync.RWMutex
}
func (d *Database) Insert(s string,i int){
	d.mutex.RLock()
	d.repo[s]=i
	d.mutex.RUnlock()
	fmt.Printf("Запись добавлена в базу! Ключ: %s, значение: %d\n",s,i)
}
func (d *Database) Select(s string) error{
	d.mutex.RLock()
	i,ok:=d.repo[s]
	d.mutex.RUnlock()
	if ok{
		fmt.Printf("Ключ: %s  значение: %d\n",s,i)
		return nil}else{
			return  errors.New("запись не найдена")}
}
type command interface {
	execute() error
}
type InsertInDatabase struct {
	database *Database
	key string
	value int
}
func (iid *InsertInDatabase) execute() error {iid.database.Insert(iid.key,iid.value)
	return nil}

type SelectFromDatabase struct {
	database *Database
	key string
	err error
}
func (sfd *SelectFromDatabase) execute() error{err:=sfd.database.Select(sfd.key)
sfd.err=err
return err
}

func main (){
	database:=&Database{repo: make(map[string]int)}
	InsertCommand:=&InsertInDatabase{database: database,key: "first",value: 1}
	SelectCommand1:=&SelectFromDatabase{database: database,key: "first"}
	SelectCommand2:=&SelectFromDatabase{database: database,key: "second"}
commands:=[...] command{InsertCommand,SelectCommand1,SelectCommand2}
for _,com:=range commands{
	err:=com.execute()
	if err!=nil{fmt.Println(err.Error())}
}
}

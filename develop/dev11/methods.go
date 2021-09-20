package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strconv"
	"time"
)

func parseConfig(value string) string {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error", err)
	}
	return os.Getenv(value)
}
func NewRepo() *repo {
	return &repo{
		myMap: make(map[string]string),
		arrayDay: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
			"15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"},
	}

}
func (r *repo) create(w http.ResponseWriter, evv string, eventT string) {
	r.myMap[evv] = eventT
	result := resultAndError{Result: "Событие создано успешно!"}
	makeJSON(w, result)
}
func (r *repo) update(w http.ResponseWriter, evv string, eventT string) {
	var result resultAndError
	_, ok := r.myMap[evv]
	if ok {
		r.myMap[evv] = eventT
		result = resultAndError{Result: "Событие обновлено успешно!"}
	} else {
		result = resultAndError{Err: "Значение не найдено!"}
	}
	makeJSON(w, result)
}
func (r *repo) delete(w http.ResponseWriter, evv string) {
	var result resultAndError
	_, ok := r.myMap[evv]
	if ok {
		delete(r.myMap, evv)
		result = resultAndError{Result: "Событие удалено успешно!"}
	} else {
		result = resultAndError{Err: "Значение не найдено!"}
	}
	makeJSON(w, result)
}
func (r *repo) getForDay(w http.ResponseWriter, evv string, day, month, year int) {
	value, ok := r.myMap[evv]
	if ok {
		newEvent := event{Day: day, Month: month, Year: year, Event: value}
		newOutput := outputDay{ResultDay: newEvent}
		makeJSON(w, newOutput)
	} else {
		result := resultAndError{Err: "Значение не найдено!"}
		makeJSON(w, result)
	}
}
func (r *repo) getForMonth(w http.ResponseWriter, month, year int) {
	var events []event
	for _, vvv := range r.arrayDay {
		value, ok := r.myMap[fmt.Sprintf("%d/%d/%s", year, month, vvv)]
		vv, _ := strconv.Atoi(vvv)
		if ok {
			newEvent := event{Day: vv, Month: month, Year: year, Event: value}
			events = append(events, newEvent)
		}
	}
	if len(events) == 0 {
		result := resultAndError{Err: "Значение не найдено!"}
		makeJSON(w, result)
		return
	}
	NewOutput := output{Result: events}
	makeJSON(w, NewOutput)
}
func (r *repo) getForWeek(w http.ResponseWriter, evv string) {
	var events []event
	layout := "2006/1/2"
	t, err := time.Parse(layout, evv)
	if err != nil {
		fmt.Printf("%v", err)
	}
	nDay := int(t.Weekday())
	if nDay == 0 {
		nDay = 7
	}
	for i := 1 - nDay; i <= 7-nDay; i++ {
		time1 := t.AddDate(0, 0, i)
		value, ok := r.myMap[fmt.Sprintf("%d/%d/%d", time1.Year(), time1.Month(), time1.Day())]
		if ok {
			newEvent := event{Day: time1.Day(), Month: int(time1.Month()), Year: time1.Year(), Event: value}
			events = append(events, newEvent)
		}
	}
	if len(events) == 0 {
		result := resultAndError{Err: "Значение не найдено!"}
		makeJSON(w, result)
		return
	}
	NewOutput := output{Result: events}
	makeJSON(w, NewOutput)

}

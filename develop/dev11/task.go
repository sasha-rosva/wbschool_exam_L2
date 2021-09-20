package main

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type event struct {
	Day   int
	Month int
	Year  int
	Event string
}
type output struct {
	Result []event `json:"result,omitempty"`
}
type outputDay struct {
	ResultDay event `json:"result,omitempty"`
}
type resultAndError struct {
	Result string `json:"result,omitempty"`
	Err    string `json:"error,omitempty"`
}
type repo struct {
	myMap    map[string]string
	arrayDay []string
}

func makeJSON(w http.ResponseWriter, i interface{}) {
	jSon, err := json.Marshal(i)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, _ = w.Write(jSon)
}

type repository interface {
	create(w http.ResponseWriter, evv string, eventT string)
	update(w http.ResponseWriter, evv string, eventT string)
	delete(w http.ResponseWriter, evv string)
	getForDay(w http.ResponseWriter, evv string, day, month, year int)
	getForWeek(w http.ResponseWriter, evv string)
	getForMonth(w http.ResponseWriter, month, year int)
}

type Handler struct {
	r repository
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	goodRequestBool := true
	var evv string
	start := time.Now()
	eventT := req.FormValue("event")
	evv = req.FormValue("year") + "/" + req.FormValue("month") + "/" + req.FormValue("day")
	day, _ := strconv.Atoi(req.FormValue("day"))
	month, _ := strconv.Atoi(req.FormValue("month"))
	year, _ := strconv.Atoi(req.FormValue("year"))
	if _, err := time.Parse("2006/1/2", evv); err != nil && req.URL.Path != "/events_for_month" {
		w.WriteHeader(400)
		return
	}
	switch req.URL.Path {
	case "/create_event":
		if req.Method != http.MethodPost {
			w.WriteHeader(405)
			return
		}
		h.r.create(w, evv, eventT)

	case "/update_event":
		if req.Method != http.MethodPost {
			w.WriteHeader(405)
			return
		}
		h.r.update(w, evv, eventT)

	case "/delete_event":
		if req.Method != http.MethodPost {
			w.WriteHeader(405)
			return
		}
		h.r.delete(w, evv)
	case "/events_for_day":
		if req.Method != http.MethodGet {
			w.WriteHeader(405)
			return
		}
		h.r.getForDay(w, evv, day, month, year)

	case "/events_for_month":
		if req.Method != http.MethodGet {
			w.WriteHeader(405)
			return
		}
		h.r.getForMonth(w, month, year)

	case "/events_for_week":
		if req.Method != http.MethodGet {
			w.WriteHeader(405)
			return
		}
		h.r.getForWeek(w, evv)
	default:
		http.NotFound(w, req)
		goodRequestBool = false
	}
	if goodRequestBool {
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	}
}

func main() {
	port := parseConfig("Port")
	repo := NewRepo()
	handler := &Handler{r: repo}
	fmt.Println("Start server!!!")
	log.Fatal(http.ListenAndServe(port, handler))

}

package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/train-do/Golang-Restfull-API/model"
	"github.com/train-do/Golang-Restfull-API/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service}
}

func (h *OrderHandler) GetAllOrder(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.GetAllOrder()
	data := struct {
		Orders []model.Order
	}{
		Orders: orders,
	}
	fmt.Printf("%+v\n", data)
	if err != nil {
		log.Fatalf("Error GetAllBook: %v", err)
	}
	// var mapFunc = make(map[string]any)
	// mapFunc["formatDate"] = func(t time.Time) string {
	// 	return t.Format("02 January 2006")
	// }
	var funcMap = template.FuncMap{
		"formatDate": func(t time.Time) string {
			return t.Format("02 January 2006")
		},
	}
	// templates = template.New("")
	// _, err = templates.Funcs(funcMap).ParseGlob("template/*.html")
	// if err != nil {
	// 	log.Fatalf("Error parsing templates: %v", err)
	// }
	// err = templates.ExecuteTemplate(w, "order-list.html", data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	tmpl, err := template.New("order-list").Funcs(funcMap).ParseFiles("template/order-list.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	err = tmpl.ExecuteTemplate(w, "order-list.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Execute template
	// err = tmpl.Execute(w, data)
	// if err != nil {
	// 	log.Fatalf("Error executing template: %v", err)
	// }
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	req := model.BodyOrder{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Error Decode:", err)
		badResponse := model.ResError{
			Status:  "error",
			Message: "Invalid request data",
			Error:   nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}
	response, err := h.service.CreateOrder(&req)
	if err != nil {
		fmt.Println("Error Query:", err)
		// badResponse := model.ResError{
		// 	Status:  "error",
		// 	Message: "Invalid request data",
		// 	Error:   nil,
		// }
		// json.NewEncoder(w).Encode(badResponse)
		// return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

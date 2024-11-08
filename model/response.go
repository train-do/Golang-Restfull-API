package model

import "time"

type Response struct {
	Status  string
	Message string
	Data    interface{}
}
type ResError struct {
	Status  string
	Message string
	Error   interface{}
}
type DataOrder struct {
	OrderCode    string
	CustomerName string
	TotalAmount  int
	FinalAmount  int
	OrderDate    time.Time
	Status       string
	OrderItem    []OrderItem
}
type OrderItem struct {
	BookCode string
	BookName string
	Qty      int
	Price    int
	SubTotal int
}

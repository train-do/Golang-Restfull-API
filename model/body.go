package model

import "time"

type BodyOrder struct {
	OrderCode       string
	CustomerName    string
	ShippingAddress interface{}
	PaymentMethod   string
	TotalAmount     int
	OrderDate       time.Time
	OrderItem       []struct {
		BookCode string
		Qty      int
	}
}

type BodyReview struct {
	Id           int
	OrderCode    string
	BookCode     string
	CustomerName string
	Rating       float32
	Review       string
	ReviewDate   time.Time
}

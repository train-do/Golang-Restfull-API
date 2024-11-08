package model

import "time"

type Order struct {
	Id          int
	OrderCode   string
	Name        string
	TotalAmount int
	FinalAmount int
	Status      bool
	CreatedAt   time.Time
}

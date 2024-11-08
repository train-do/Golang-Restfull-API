package model

import "time"

type Review struct {
	Id         int
	OrderCode  string
	BookCode   string
	Rating     float32
	Review     string
	ReviewDate time.Time
}

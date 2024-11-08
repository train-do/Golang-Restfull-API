package handler

import (
	"net/http"

	"github.com/train-do/Golang-Restfull-API/service"
)

type ReviewHandler struct {
	service *service.ReviewService
}

func NewReviewHandler(service *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{service}
}

func (h *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {}

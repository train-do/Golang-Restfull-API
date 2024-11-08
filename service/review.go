package service

import (
	"github.com/train-do/Golang-Restfull-API/model"
	"github.com/train-do/Golang-Restfull-API/repository"
)

type ReviewService struct {
	repo *repository.ReviewRepository
}

func NewReviewService(repo *repository.ReviewRepository) *ReviewService {
	return &ReviewService{repo}
}

func (s *ReviewService) CreateReview(product *model.Review) error {
	return s.repo.Create(product)
}

package router

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/train-do/Golang-Restfull-API/database"
	"github.com/train-do/Golang-Restfull-API/handler"
	mid "github.com/train-do/Golang-Restfull-API/middleware"
	"github.com/train-do/Golang-Restfull-API/repository"
	"github.com/train-do/Golang-Restfull-API/service"
)

func NewRouter() *chi.Mux {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	hUser := handler.NewUserHandler()
	rBook := repository.NewBookRepository(db)
	sBook := service.NewBookService(rBook)
	hBook := handler.NewBookHandler(sBook)
	rOrder := repository.NewOrderRepository(db)
	sOrder := service.NewOrderService(rOrder)
	hOrder := handler.NewOrderHandler(sOrder)
	rReview := repository.NewReviewRepository(db)
	sReview := service.NewReviewService(rReview)
	hReview := handler.NewReviewHandler(sReview)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/login", hUser.Login)
	router.Post("/login", hUser.Login)
	router.Group(func(r chi.Router) {
		r.Use(mid.Authentication)
		r.Get("/dashboard", hBook.Dashboard)
		r.Get("/addBook", hBook.CreateBook)
		r.Post("/addBook", hBook.CreateBook)
		r.Get("/books", hBook.GetAllBook)
		r.Get("/orders", hOrder.GetAllOrder)
		r.Put("/discount/{id}", hBook.UpdateBook)
		r.Get("/logout", hUser.Logout)
		r.Post("/logout", hUser.Logout)
	})
	router.Group(func(r chi.Router) {
		r.Use(mid.Authentication)
		r.Get("/dashboard", hBook.Dashboard)
		r.Post("/addBook", hBook.CreateBook)
		r.Get("/books", hBook.GetAllBook)
		r.Get("/orders", hOrder.GetAllOrder)
		r.Put("/discount", hBook.UpdateBook)
		r.Get("/logout", hUser.Logout)
		r.Post("/logout", hUser.Logout)
	})
	router.Group(func(r chi.Router) {
		r.Get("/customer/order", hOrder.CreateOrder)
		r.Post("/customer/review", hReview.CreateReview)
	})

	return router
}

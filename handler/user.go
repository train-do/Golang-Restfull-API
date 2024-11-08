package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

var Token string
var templates *template.Template

type UserHandle struct{}

func NewUserHandler() *UserHandle {
	return &UserHandle{}
}

func (h *UserHandle) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("MASUK GET LOGIN HANDLER")
		templates, err := template.ParseGlob("template/*.html")
		if err != nil {
			log.Fatalf("Error parsing templates: %v", err)
		}
		err = templates.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPost {
		fmt.Println("MASUK POST LOGIN HANDLER")
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username != "admin" && password != "123456" {
			unauthorized := struct {
				StatusCode int
				Message    string
			}{
				StatusCode: http.StatusUnauthorized,
				Message:    "Login Failed",
			}
			json.NewEncoder(w).Encode(unauthorized)
			return
		}
		Token = uuid.New().String()
		cookie := http.Cookie{
			Name:   "token",
			Value:  Token,
			Path:   "/",
			Domain: "localhost",
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}
func (h *UserHandle) Logout(w http.ResponseWriter, r *http.Request) {
}

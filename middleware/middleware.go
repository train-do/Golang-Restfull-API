package middleware

import (
	"net/http"

	"github.com/train-do/Golang-Restfull-API/handler"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || err == http.ErrNoCookie {
			// fmt.Println("NO COOKIE")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if cookie.Value != handler.Token {
			// fmt.Println("GAGAL VALIDASI")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

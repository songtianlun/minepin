package handle

import (
	"minepin/model"
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := model.CheckSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		next.ServeHTTP(w, r)
	}
}

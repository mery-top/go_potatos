package middleware

import(
	"go_potatos/utils"
	"net/http"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		cookie, err := r.Cookie("token")
		if err !=nil{
			http.Redirect(w, r,	"/login", http.StatusSeeOther)
			return
		}

		_, err:= ValidateJWT(cookie.Value)

		if err!=nil{
			http.Redirect(w,r, "/login", http.StatusSeeOther)
			return
		}


		next(w,r)
}

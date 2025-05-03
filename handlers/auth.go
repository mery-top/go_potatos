package handlers

import(
	"go_potatos/models"
	"go_potatos/utils"
	"http/template"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request){
	if r.Method = http.MethodGet{
		t, _:= template.ParseFiles("templates/login.html")
		t.execute(w, nil)
		return
	}

	email:= r.FormValue("title")
	password:= r.FormValue("password")

	user, err:= models.GetUserByEmail(email)
	if err !=nil || user.Password != password{
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _:= utils.GenerateJWT(user.Email)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: token,
		HttpOnly: true,
	})

	http.Redirect(w,r, "/dashboard", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request){
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: "",
		MaxAge: -1,
	})

	http.Redirect(w,r, "/login", http.StatusSeeOther)
}

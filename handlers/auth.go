package handlers

import(
	"go_potatos/models"
	"go_potatos/utils"
	"http/template"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		t, _:= template.ParseFiles("templates/login.html")
		t.execute(w, nil)
		return
	}

	email:= r.FormValue("title")
	password:= r.FormValue("password")

	user, err:= models.GetUserByEmail(email)
	if err != nil {
    http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.password, []byte(password)))

	if err!=nil{
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


func RegisterPage(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		t, _:= template.ParseFiles("templates/register.html")
		t.execute(w, nil)
		return
	}

	username:= r.FormValue("username")
	email:= r.FormValue("email")
	password:= r.FormValue("password")


	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err!=nil{
		http.Error(w, "Hashing went wrong", http.StatusInternalServerError)
		return
	}

	err= models.CreateUser(username, email, string(hashedPassword))

	if err!=nil{
		http.Error(w, "Error in Creating User", http.StatusInternalServerError)
		return
	}

	http.Redirect(r,w, "/login", http.StatusSeeOther)
}

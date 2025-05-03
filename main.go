package main

import(
	"go_potatos/database"
	"go_potatos/handlers"
	"go_potatos/middleware"
	"fmt"
	"net/http"
)


func main(){
	database.InitDB()
	http.Handle("/assests/", http.StripPrefix("/assests", http.FileServer(http.Dir("assests")
	http.HandleFunc("/login", handlers.LoginPage)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/register", handlers.RegisterPage)
	http.HandlerFunc("/dashboard", middleware.RequireAuth(handlers.DashboardHandler))

	http.HandlerFunc("/create", middleware.RequireAuth(handlers.CreatePostPage))

	fmt.Println("Server running at PORT 8080")
	http.ListenAndServe(":8080", nil)

}

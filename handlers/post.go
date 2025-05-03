package handlers

import(
	"go_potatos/models"
	"go_potatos/utils"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	)

func CreatePostPage(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		t, _:= template.ParseFiles("templates/create_post.html")
		t.Execute(w, nil)
		return
	}

	r.ParseMultiPartForm(10 << 20)
	title:= r.FormValue("title")
	content:= r.FormValue("content")
	file, handler,err := r.FormFile("image")

	var imagePath string

	if err!=nil{
		defer file.Close()
		imagePath = filepath.Join("assests", handler.Filename)
		dst, err:= os.Create(imagePath)
		
		if err !=nil{
			http.Error(w, "Unable to Open", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err:= dst.ReadFrom(file)

		if err !=nil{
			http.Error(w, "Failed to Read File", http.StatusInternalServerError)
			return
		}
	}
	//Check Cookies

	cookie, err:= r.Cookie("token")
	if err!=nil{
		http.Redirect(w,r, "/login", http.StatusInternalServerError)
		return
	}

	claims, err:= utils.ValidateJWT(cookie.Value)
	if err!=nil{
		http.Redirect(w,r, "/login", http.StatusInternalServerError)
		return
	}

	user, err:= models.GetUserByEmail(claims.Email)
	if err!=nil{
		http.Redirect(w,r, "/login", http.StatusInternalServerError)
		return
	}

	post:= &models.Post{
		Title: title,
		Content: content,
		ImagePath: imagePath,
		AuthorID: user.ID
	}

	err:= models.CreatePost(post)

	if err!=nil{
		http.Error(w, "Can't Create Posts", http.StatusInternalServerError)
		return
	}

	http.Redirect(w,r , "/dashboard", http.StatusSeeOther)
}

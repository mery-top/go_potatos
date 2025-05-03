package models

import(
	"go_potatos/database"
	"time"
)

type Post struct{
	ID int
	Title string
	Content string
	ImagePath string
	AuthorID int
	CreatedAt time.Time
}

func CreatePost(post *Post) error{
	_,err:=database.db.Exec("insert into posts(title, content, image_path, author_id) values($1, $2, $3, $4)", post.Title, post.Content, post.ImagePath, post.AuthorID)
	return nil
}

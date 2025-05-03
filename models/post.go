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


func GetPostByUserID(db *sql.DB, userID int) ([]Post, error){C
	rows, err:= db.Query("select id, title, content, image_path, created_at where author_id = $1 orderby created_at desc", userID)
	if err!=nil{
		return nil, err
	}

	defer rows.Close()
	var posts []Post
	for rows.Next(){
		var post Post
		err:= rows.Scan(&post.ID, &post.Title, &post.Content, &post.ImagePath, &post.AuthorID, &post.CreatedAt)
		if err!=nil{
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

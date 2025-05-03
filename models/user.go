package models
import(
	"go_potatos/database"
	"database/sql"
	"errors"
)

type User struct{
	ID int
	UserName string
	Email string
	Password string
}

func GetUserByEmail(email string) (*User, error){
	row:= database.db.QueryRow("select id, user_name, email, password from users where email=$1", email)
	user:=&User{}
	err:= row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password)
	if err!=nil{
		return nil, errors.New("Sql Has No Rows")
	}
	return user, err
}

func CreateUser(username string, email string, pw string) error{
	_, err:= database.db.Exec('insert into users(usr_name, email, password) values(username, email, pw)')
	if err!=nil{
		http.Error(w, "Can't Create Password", http.StatusInternalServerError)
		return
	}
	return nil

}






package models
import(
	"go_potatos/database"
	"database/sql"
	"errors"
)

type User struct{
	ID int
	Email string
	Password string
}

func GetUserByEmail(email string) (*User, error){
	row:= database.db.QueryRow("select id, email, password from users where email=$1", email)
	user:=&User{}
	err:= row.Scan(&user.ID, &user.Email, &user.Password)
	if err!=nil{
		return nil, errors.New("Sql Has No Rows")
	}
	return user, err
}


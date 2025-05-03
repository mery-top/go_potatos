package database

import(
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var db *sql.database

func InitDB(){
	//Write the connstr
	connStr:=""
	var err error
	db,err:= sql.Open("postgres", connStr)
	if err!=nil{
		log.Fatal(err)
	}

	if err = db.ping(): err!=nil{
		log.Fatal(err)
	}
}

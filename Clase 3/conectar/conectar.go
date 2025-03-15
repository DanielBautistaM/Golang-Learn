package conectar

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// funcion para conectar a la db
var Db *sql.DB

func Conectar() {
	errorVariable := godotenv.Load()
	if errorVariable != nil {
		panic(errorVariable)
	}
	conection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}
	Db = conection
}

func CerrarConexion() {
	Db.Close()
}

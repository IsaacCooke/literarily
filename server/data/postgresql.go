package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBInfo struct {
  host string
  port string
  user string
  password string
  dbname string
}

func populateDBInfo() DBInfo {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  return DBInfo {
    host: os.Getenv("HOST"),
    port: os.Getenv("PORT"),
    user: os.Getenv("USER"),
    password: os.Getenv("PASSWORD"),
    dbname: os.Getenv("DBNAME"),
  }
}

func Connect() *sql.DB {
  dbInfo := populateDBInfo()

  port, err := strconv.Atoi(dbInfo.port)
  checkError(err)

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable", dbInfo.host, port, dbInfo.user, dbInfo.password, dbInfo.dbname)

  db, err := sql.Open("postgres", psqlInfo)
  checkError(err)

  defer db.Close()

  err = db.Ping()
  checkError(err)

  fmt.Println("Connected to database")
  return db
}

func checkError(err error){
  if err != nil {
    panic(err)
  }
}

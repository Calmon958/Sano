package main

import (
	"Sano/database"
	_ "github.com/mattn/go-sqlite3"

)

func main(){
	db.InitDB()
	db.CreateTable()
}
package main

import (
	"gin-todo/db"
	"gin-todo/server"
)

func main() {
	db.InitDB()
	server.Init()
	defer db.CloseDB()

}

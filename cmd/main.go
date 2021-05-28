package main

import (
	"log"

	"github.com/h-u-m-a-n/todo-app"
	"github.com/h-u-m-a-n/todo-app/package/handler"
)


func main() {
	handlers := new(handler.Handler)
	
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil{
		log.Fatalf("error occuring while running http server: %s", err.Error())
	}
}
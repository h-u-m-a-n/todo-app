package main

import (
	"log"

	"github.com/h-u-m-a-n/todo-app"
	"github.com/h-u-m-a-n/todo-app/package/handler"
	"github.com/h-u-m-a-n/todo-app/package/repository"
	"github.com/h-u-m-a-n/todo-app/package/service"
	"github.com/spf13/viper"
)


func main() {
	if err := initConfig(); err != nil{
		log.Fatalf("error initialization configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil{
		log.Fatalf("error occuring while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("configs/config.yml")
	return viper.ReadInConfig()
}
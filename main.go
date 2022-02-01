package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mrrizal/todo-api-clean-arch/api"
	"github.com/mrrizal/todo-api-clean-arch/config"
	"github.com/mrrizal/todo-api-clean-arch/repository/mysql"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to log .env file.")
	}
}

func main() {
	var cfg config.Config

	loadEnv()
	config.GetConfig(&cfg)

	_, err := mysql.NewMysqlDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	api.SetupRoutes(app)
	app.Listen(":5000")
}

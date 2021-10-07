package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/smokers10/go-diagem.git/infrastructure/database"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
	"github.com/smokers10/go-diagem.git/router/api"
)

func main() {
	engine := html.New("./view", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// koneksi database
	mysql, err := database.MYSQLConn(os.Getenv("PRODUCTION_MODE"))
	if err != nil {
		panic(err)
	}

	app.Use(recover.New())

	app.Static("/", "./public")

	serviceResolver := resolver.MYSQLResolver(mysql)

	api.AdminAPI(app, &serviceResolver)

	log.Fatal(app.Listen(":8001"))
}

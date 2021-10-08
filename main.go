package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/smokers10/go-diagem.git/infrastructure/database"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
	"github.com/smokers10/go-diagem.git/router/api"
)

func main() {
	// konfigurasi view engine
	engine := html.New("./view", ".html")

	// konfigurasi & init fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// koneksi database
	mysql, err := database.MYSQLConn(os.Getenv("PRODUCTION_MODE"))
	if err != nil {
		panic(err)
	}

	// error recovery
	// app.Use(recover.New())

	// static folder
	app.Static("/", "./public")

	// resolver init
	serviceResolver := resolver.MYSQLResolver(mysql)

	// router
	api.AdminAPI(app, &serviceResolver)
	api.UserAPI(app, &serviceResolver)

	// serving
	log.Fatal(app.Listen(":8001"))
}

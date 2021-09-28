package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/smokers10/go-diagem.git/infrastructure/database"
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
	fmt.Println(mysql)

	app.Use(recover.New())

	app.Static("/", "./public")

	log.Fatal(app.Listen(":8000"))
}

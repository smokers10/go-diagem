package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/smokers10/go-diagem.git/infrastructure/config"
	"github.com/smokers10/go-diagem.git/infrastructure/database"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
	"github.com/smokers10/go-diagem.git/infrastructure/session"
	"github.com/smokers10/go-diagem.git/router/api"
	"github.com/smokers10/go-diagem.git/router/web"
)

func main() {
	fmt.Println(config.ReadConfig())
	config := config.ReadConfig().Application

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static folder
	app.Static("/", "./public")

	// error recovery
	if mode := config.APP_Production_Mode; mode == "production" {
		app.Use(recover.New())
	}

	// koneksi database
	mysql, err := database.MYSQLConn()
	if err != nil {
		panic(err)
	}

	// send email test

	// resolver init
	serviceResolver := resolver.MYSQLResolver(mysql)

	// session init
	session := session.MYSQLSessionStore(os.Getenv("PRODUCTION_MODE"))

	// router
	api.AdminAPI(app, &serviceResolver)
	api.UserAPI(app, &serviceResolver)
	web.AdminWebPage(app, session, &serviceResolver)
	web.UserWebPage(app, session, &serviceResolver)

	// serving
	log.Fatal(app.Listen(config.APP_PORT))
}

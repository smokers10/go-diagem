package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/smokers10/go-diagem.git/infrastructure/database"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
	"github.com/smokers10/go-diagem.git/infrastructure/session"
	"github.com/smokers10/go-diagem.git/router/api"
	"github.com/smokers10/go-diagem.git/router/web"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static folder
	app.Static("/", "./public")

	// error recovery
	// app.Use(recover.New())

	// koneksi database
	mysql, err := database.MYSQLConn(os.Getenv("PRODUCTION_MODE"))
	if err != nil {
		panic(err)
	}

	// resolver init
	serviceResolver := resolver.MYSQLResolver(mysql)

	// session init
	session := session.MYSQLSessionStore(os.Getenv("PRODUCTION_MODE"))

	// router
	api.AdminAPI(app, &serviceResolver)
	api.UserAPI(app, &serviceResolver)
	web.AdminWebPage(app, session, &serviceResolver)
	web.UserWebPage(app, session, &serviceResolver)

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Render("test-page", nil)
	})

	// serving
	log.Fatal(app.Listen(":8001"))
}

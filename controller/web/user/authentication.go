package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/smokers10/go-diagem.git/domain"
)

type authenticationController struct {
	userService domain.UserService
	session     session.Store
}

func AuthenticationController(user *domain.UserService, session *session.Store) authenticationController {
	return authenticationController{
		userService: *user,
		session:     *session,
	}
}

func (a *authenticationController) LoginPage(c *fiber.Ctx) error {
	return c.Render("umum/auth/login", nil)
}

func (a *authenticationController) RegisterPage(c *fiber.Ctx) error {
	return c.Render("umum/auth/register", nil)
}

func (a *authenticationController) Login(c *fiber.Ctx) error {
	credential := domain.UserCredential{}
	sess, err := a.session.Get(c)
	if err != nil {
		panic(err)
	}

	if err := c.BodyParser(&credential); err != nil {
		panic(err)
	}

	response := a.userService.Login(&credential)

	if response.Success {
		sess.Set("token", fmt.Sprintf("Bearer %s", response.Token))
		sess.Save()
	}

	return c.Status(response.Status).JSON(response)
}

func (a *authenticationController) Register(c *fiber.Ctx) error {
	basicData := domain.UserBasicData{}
	sess, err := a.session.Get(c)
	if err != nil {
		panic(err)
	}

	if err := c.BodyParser(&basicData); err != nil {
		panic(err)
	}

	response := a.userService.Registrasi(&basicData)

	if response.Success {
		sess.Set("token", fmt.Sprintf("Bearer %s", response.Token))
		sess.Save()
	}

	return c.Status(response.Status).JSON(response)
}

func (a *authenticationController) Logout(c *fiber.Ctx) error {
	sess, err := a.session.Get(c)
	if err != nil {
		panic(c)
	}

	sess.Destroy()

	return c.Redirect("/")
}

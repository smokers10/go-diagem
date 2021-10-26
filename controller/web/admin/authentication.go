package admin

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/smokers10/go-diagem.git/domain"
)

type adminAuthenticationController struct {
	adminService domain.AdminService
	session      session.Store
}

func AdminAuthenticationController(admin *domain.AdminService, session *session.Store) adminAuthenticationController {
	return adminAuthenticationController{
		adminService: *admin,
		session:      *session,
	}
}

func (aac *adminAuthenticationController) LoginPage(c *fiber.Ctx) error {
	return c.Render("admin/auth/login", nil)
}

func (aac *adminAuthenticationController) Login(c *fiber.Ctx) error {
	cred := domain.AdminCredentials{}
	sess, err := aac.session.Get(c)
	if err != nil {
		panic(err)
	}

	if err := c.BodyParser(&cred); err != nil {
		panic(err)
	}

	response := aac.adminService.Login(&cred)

	if response.Success {
		sess.Set("token", fmt.Sprintf("Bearer %s", response.Token))
		sess.Save()
	}

	return c.Status(response.Status).JSON(response)
}

func (aac *adminAuthenticationController) Logout(c *fiber.Ctx) error {
	sess, err := aac.session.Get(c)
	if err != nil {
		panic(err)
	}

	sess.Destroy()

	return c.Redirect("/admin/login")
}

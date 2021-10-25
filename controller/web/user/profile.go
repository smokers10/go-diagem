package user

import "github.com/gofiber/fiber/v2"

type userProfileController struct{}

func UserProfileCOntroller() userProfileController {
	return userProfileController{}
}

func (u *userProfileController) ProfilePage(c *fiber.Ctx) error {
	return c.Render("umum/user/profile", nil)
}

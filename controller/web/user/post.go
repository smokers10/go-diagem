package user

import "github.com/gofiber/fiber/v2"

type postController struct{}

func PostController() postController {
	return postController{}
}

func (p *postController) IndexPage(c *fiber.Ctx) error {
	return c.Render("umum/post/index", nil)
}

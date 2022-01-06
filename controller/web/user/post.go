package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
)

type postController struct {
	blogService domain.BlogService
}

func PostController(blog *domain.BlogService) *postController {
	return &postController{
		blogService: *blog,
	}
}

func (p *postController) IndexPage(c *fiber.Ctx) error {
	return c.Render("umum/post/index", fiber.Map{"Logged": etc.GetLocal(c)})
}

func (p *postController) ReadBlogPage(c *fiber.Ctx) error {
	return c.Render("umum/post/read", fiber.Map{"Logged": etc.GetLocal(c)})
}

func (p *postController) GetAll(c *fiber.Ctx) error {
	judul := c.FormValue("judul")
	response := p.blogService.PublishedOnly(judul)
	return c.Status(response.Status).JSON(response)
}

func (p *postController) ReadBlog(c *fiber.Ctx) error {
	response := p.blogService.Detail(c.Params("slug"))
	return c.Status(response.Status).JSON(response)
}

func (p *postController) GetLatest(c *fiber.Ctx) error {
	response := p.blogService.Latest()
	return c.Status(response.Status).JSON(response)
}

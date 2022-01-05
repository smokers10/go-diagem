package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type blogController struct {
	blogService domain.BlogService
}

func BlogController(blog *domain.BlogService) *blogController {
	return &blogController{blogService: *blog}
}

func (bc *blogController) IndexPage(c *fiber.Ctx) error {
	return c.Render("admin/blog/index", nil)
}

func (bc *blogController) TambahPage(c *fiber.Ctx) error {
	return c.Render("admin/blog/tambah", nil)
}

func (bc *blogController) EditPage(c *fiber.Ctx) error {
	return c.Render("admin/blog/edit", nil)
}

func (bc *blogController) Create(c *fiber.Ctx) error {
	req := domain.Blog{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := bc.blogService.Create(&req)

	return c.Status(response.Status).JSON(response)
}

func (bc *blogController) Update(c *fiber.Ctx) error {
	req := domain.Blog{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := bc.blogService.Update(&req)

	return c.Status(response.Status).JSON(response)
}

func (bc *blogController) Delete(c *fiber.Ctx) error {
	req := domain.Blog{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := bc.blogService.Delete(req.ID)

	return c.Status(response.Status).JSON(response)
}

func (bc *blogController) ReadAll(c *fiber.Ctx) error {
	response := bc.blogService.ReadAll()

	return c.Status(response.Status).JSON(response)
}

func (bc *blogController) Detail(c *fiber.Ctx) error {
	response := bc.blogService.Detail(c.Params("slug"))

	return c.Status(response.Status).JSON(response)
}

func (bc *blogController) UpdateThumbnail(c *fiber.Ctx) error {
	req := domain.Blog{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := bc.blogService.UpdateThumbnail(&req)

	return c.Status(response.Status).JSON(response)
}

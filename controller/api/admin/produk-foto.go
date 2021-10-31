package admin

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type produkFotoController struct {
	produkFotoService *domain.ProdukFotoService
}

func ProdukFotoController(produkFoto *domain.ProdukFotoService) produkFotoController {
	return produkFotoController{produkFotoService: produkFoto}
}

func (p *produkFotoController) Create(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}

	files := form.File["image"]

	for _, file := range files {
		fmt.Println(file.Filename, file.Header, file.Size)
	}

	return nil
}

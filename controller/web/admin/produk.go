package admin

import "github.com/gofiber/fiber/v2"

type produkController struct{}

func ProdukController() produkController {
	return produkController{}
}

func (p *produkController) IndexPage(c *fiber.Ctx) error {
	return c.Render("admin/produk/index", nil)
}

func (p *produkController) TambahPage(c *fiber.Ctx) error {
	return c.Render("admin/produk/tambah", nil)
}

func (p *produkController) EditPage(c *fiber.Ctx) error {
	return c.Render("admin/produk/edit", nil)
}

func (p *produkController) KategoriPage(c *fiber.Ctx) error {
	return c.Render("admin/produk/kategori", nil)
}

func (p *produkController) KategoriTambahPage(c *fiber.Ctx) error {
	return c.Render("admin/produk/kategori-form", nil)
}

func (p *produkController) KategoriEditPage(c *fiber.Ctx) error {
	return c.Render("admin/produk/kategori-edit", nil)
}

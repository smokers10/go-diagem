package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type feedbackController struct {
	feedbackSrv domain.FeedbackService
}

func FeedbackController(feedback *domain.FeedbackService) *feedbackController {
	return &feedbackController{feedbackSrv: *feedback}
}

func (fc *feedbackController) GiveFeedback(c *fiber.Ctx) error {
	req := domain.Feedback{}
	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}
	response := fc.feedbackSrv.Create(&req)
	return c.Status(response.Status).JSON(response)
}

func (fc *feedbackController) EditFeedback(c *fiber.Ctx) error {
	req := domain.Feedback{}
	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}
	response := fc.feedbackSrv.Update(&req)
	return c.Status(response.Status).JSON(response)
}

func (fc *feedbackController) Delete(c *fiber.Ctx) error {
	req := domain.Feedback{}
	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}
	response := fc.feedbackSrv.Delete(&req)
	return c.Status(response.Status).JSON(response)
}

func (fc *feedbackController) Read(c *fiber.Ctx) error {
	produkID := c.Params("produk_id")
	response := fc.feedbackSrv.Read(produkID)
	return c.Status(response.Status).JSON(response)
}

func (fc *feedbackController) GetMyFeedback(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	response := fc.feedbackSrv.ByUserID(userID)
	return c.Status(response.Status).JSON(response)
}

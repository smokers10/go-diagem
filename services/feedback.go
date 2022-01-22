package services

import (
	"fmt"

	"github.com/smokers10/go-diagem.git/domain"
)

type feedbackServiceImpl struct {
	feedbackRepository domain.FeedbackRepository
}

func FeedbackService(feedback *domain.FeedbackRepository) domain.FeedbackService {
	return &feedbackServiceImpl{feedbackRepository: *feedback}
}

func (fs *feedbackServiceImpl) Create(req *domain.Feedback) *domain.Response {
	res := domain.Response{}

	// check commentar
	checkFeedback, err := fs.feedbackRepository.OnProduct(req.UserID, req.ProdukID)
	if err != nil {
		return &domain.Response{
			Message: "error saat check feedback",
			Status:  500,
		}
	}

	if checkFeedback.UserID != 0 {
		return &domain.Response{
			Message: "Hanya bisa membuat feedback satu kali",
			Status:  200,
		}
	}

	// buat feedback
	if err := fs.feedbackRepository.Create(req); err != nil {
		fmt.Println(err)
		res.Message = "Error saat menyimpan data feedback"
		res.Status = 500
		return &res
	}

	res.Message = "feeback berhasil disimpan"
	res.Status = 200
	res.Success = true
	return &res
}

func (fs *feedbackServiceImpl) Update(req *domain.Feedback) *domain.Response {
	res := domain.Response{}
	if err := fs.feedbackRepository.Update(req); err != nil {
		fmt.Println(err)
		res.Message = "Error saat menyimpan data feedback"
		res.Status = 500
		return &res
	}
	res.Message = "feeback berhasil disimpan"
	res.Status = 200
	res.Success = true
	return &res
}

func (fs *feedbackServiceImpl) Read(produkID string) *domain.Response {
	feedback, err := fs.feedbackRepository.Read(produkID)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat mengambil data feedback",
			Status:  500,
		}
	}

	// perhitungan rata-rata rating
	// dasar litelatur : https://calculator.academy/average-rating-calculator-star-rating/#f1p1|f2p0
	// R total jumlah rating
	// AR rata-rata rating
	R := float32(feedback.ByRate.OneStar + feedback.ByRate.TwoStar + feedback.ByRate.ThreeStar + feedback.ByRate.FourStar + feedback.ByRate.FiveStar)
	AR := (1*float32(feedback.ByRate.OneStar) + 2*float32(feedback.ByRate.TwoStar) + 3*float32(feedback.ByRate.ThreeStar) + 4*float32(feedback.ByRate.FourStar) + 5*float32(feedback.ByRate.FiveStar)) / R
	feedback.AverageRating = AR

	return &domain.Response{
		Data:    feedback,
		Message: "feeback berhasil diambil",
		Status:  200,
		Success: true,
	}
}

func (fs *feedbackServiceImpl) Delete(req *domain.Feedback) *domain.Response {
	res := domain.Response{}
	if err := fs.feedbackRepository.Delete(req); err != nil {
		fmt.Println(err)
		res.Message = "Error saat menyimpan data feedback"
		res.Status = 500
		return &res
	}
	res.Message = "feeback berhasil disimpan"
	res.Status = 200
	res.Success = true
	return &res
}

func (fs *feedbackServiceImpl) ByRating(productID string, rating int) *domain.Response {
	feedback, err := fs.feedbackRepository.ByRating(productID, rating)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat mengambil data feedback",
			Status:  500,
		}
	}

	return &domain.Response{
		Message: "data feedback berhasil diambil",
		Status:  200,
		Success: true,
		Data:    feedback,
	}
}

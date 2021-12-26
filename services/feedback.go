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
	res := domain.Response{}
	feedback, err := fs.feedbackRepository.Read(produkID)
	if err != nil {
		fmt.Println(err)
		res.Message = "Error saat menyimpan data feedback"
		res.Status = 500
		return &res
	}
	res.Data = feedback
	res.Message = "feeback berhasil disimpan"
	res.Status = 200
	res.Success = true
	return &res
}

func (fs *feedbackServiceImpl) ByUserID(userID int) *domain.Response {
	res := domain.Response{}
	feedback, err := fs.feedbackRepository.ByUserId(userID)
	if err != nil {
		fmt.Println(err)
		res.Message = "Error saat menyimpan data feedback"
		res.Status = 500
		return &res
	}
	res.Data = feedback
	res.Message = "feeback berhasil disimpan"
	res.Status = 200
	res.Success = true
	return &res
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

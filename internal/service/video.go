package service

import (
	"github.com/Madou-Shinni/gin-quickstart/internal/data"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/Madou-Shinni/go-logger"
	"go.uber.org/zap"
)

// 定义接口
type VideoRepo interface {
	Create(video domain.Video) error
	Delete(video domain.Video) error
	Update(video map[string]interface{}) error
	Find(video domain.Video) (domain.Video, error)
	List(page domain.PageVideoSearch) ([]domain.Video, int64, error)
	Count(page domain.PageVideoSearch) (int64, error)
	DeleteByIds(ids request.Ids) error
}

type VideoService struct {
	repo VideoRepo
}

func NewVideoService() *VideoService {
	return &VideoService{repo: &data.VideoRepo{}}
}

func (s *VideoService) Add(video domain.Video) error {
	// 3.持久化入库
	if err := s.repo.Create(video); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(video)", zap.Error(err), zap.Any("domain.Video", video))
		return err
	}

	return nil
}

func (s *VideoService) Delete(video domain.Video) error {
	if err := s.repo.Delete(video); err != nil {
		logger.Error("s.repo.Delete(video)", zap.Error(err), zap.Any("domain.Video", video))
		return err
	}

	return nil
}

func (s *VideoService) Update(video map[string]interface{}) error {
	if err := s.repo.Update(video); err != nil {
		logger.Error("s.repo.Update(video)", zap.Error(err), zap.Any("domain.Video", video))
		return err
	}

	return nil
}

func (s *VideoService) Find(video domain.Video) (domain.Video, error) {
	res, err := s.repo.Find(video)

	if err != nil {
		logger.Error("s.repo.Find(video)", zap.Error(err), zap.Any("domain.Video", video))
		return res, err
	}

	return res, nil
}

func (s *VideoService) List(page domain.PageVideoSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, total, err := s.repo.List(page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageVideoSearch", page))
		return pageRes, err
	}

	pageRes.List = data
	pageRes.Total = total

	return pageRes, nil
}

func (s *VideoService) DeleteByIds(ids request.Ids) error {
	if err := s.repo.DeleteByIds(ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}

package data

import (
	"errors"
	"fmt"
	"github.com/Madou-Shinni/gin-quickstart/common"
	glob "github.com/Madou-Shinni/gin-quickstart/global"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/pagelimit"
)

type VideoRepo struct {
}

func (s *VideoRepo) Create(video domain.Video) error {
	return global.DB.Create(&video).Error
}

func (s *VideoRepo) Delete(video domain.Video) error {
	return global.DB.Delete(&video).Error
}

func (s *VideoRepo) DeleteByIds(ids request.Ids) error {
	return global.DB.Delete(&[]domain.Video{}, ids.Ids).Error
}

func (s *VideoRepo) Update(video map[string]interface{}) error {
	var columns []string
	for key := range video {
		columns = append(columns, key)
	}
	if _, ok := video["id"]; !ok {
		// 不存在id
		return errors.New(fmt.Sprintf("missing %s.id", "video"))
	}
	model := domain.Video{}
	//model.ID = uint(video["id"].(float64))
	return global.DB.Model(&model).Select(columns).Updates(&video).Error
}

func (s *VideoRepo) Find(video domain.Video) (domain.Video, error) {
	db := global.DB.Model(&domain.Video{})
	// TODO：条件过滤

	res := db.First(&video)

	return video, res.Error
}

func (s *VideoRepo) List(page domain.PageVideoSearch) ([]domain.Video, int64, error) {
	var (
		videoList []domain.Video
		total     int64
		err       error
	)

	// page
	from, size := pagelimit.OffsetLimit(page.PageNum, page.PageSize)

	searchFields := []string{"title", "introduction"}

	videoList, total, err = common.MatchQuery[domain.Video](glob.Es, domain.Video{}.Index(), from, size, searchFields, page.Keyword)
	if err != nil {
		return nil, 0, err
	}

	return videoList, total, err
}

func (s *VideoRepo) Count(page domain.PageVideoSearch) (int64, error) {
	var (
		count int64
		err   error
	)

	db := global.DB.Model(&domain.Video{})

	// TODO：条件过滤

	err = db.Model(&domain.Video{}).Count(&count).Error

	return count, err
}

package service

import (
	"context"
	"fmt"
	glob "github.com/Madou-Shinni/gin-quickstart/global"
	"github.com/Madou-Shinni/gin-quickstart/internal/data"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain/req"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain/resp"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/Madou-Shinni/go-logger"
	"github.com/goccy/go-json"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 如果您需要将星期几转换为中文，您可以使用一个 map 来进行转换
var weekdayMap = map[time.Weekday]string{
	time.Sunday:    "周日",
	time.Monday:    "周一",
	time.Tuesday:   "周二",
	time.Wednesday: "周三",
	time.Thursday:  "周四",
	time.Friday:    "周五",
	time.Saturday:  "周六",
}

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

	res.Director = strings.TrimSuffix(res.Director, "/")
	res.Screenwriter = strings.TrimSuffix(res.Screenwriter, "/")

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

func (s *VideoService) Home() (resp.Home, error) {
	var (
		home resp.Home
	)

	query := elastic.NewMatchAllQuery()
	res, err := glob.Es.Search().Index().Query(query).Do(context.Background())
	if err != nil {
		return home, err
	}
	if res.Hits == nil {
		return home, fmt.Errorf("searchResult.Hits is nil")
	}

	for _, hit := range res.Hits.Hits {
		err = json.Unmarshal(hit.Source, &home)
		if err != nil {
			return home, err
		}
	}

	/*
		以下查询是为了获取视频的更新时间，然后转换为星期几
	*/
	// 获取ids
	var ids []interface{}
	if home.Modules != nil {
		for _, v := range home.Modules[0].ContentList {
			ids = append(ids, v.ID)
		}
	}
	// 创建 Terms 查询 查询ids中的数据
	type UpdateAtData struct {
		ID       int       `json:"id"`
		UpdateAt time.Time `json:"updateAt"`
	}
	termsQuery := elastic.NewTermsQuery("id", ids...)
	resp, err := glob.Es.Search().Index("videos").Query(termsQuery).Size(len(ids)).Do(context.Background())
	if err != nil {
		return home, err
	}

	// 获取周几
	weekMap := make(map[int]string, len(resp.Hits.Hits))
	for _, hit := range resp.Hits.Hits {
		var updateAtData UpdateAtData
		err = json.Unmarshal(hit.Source, &updateAtData)
		if err != nil {
			return home, err
		}
		// 转周几
		key := updateAtData.UpdateAt.Weekday()
		weekMap[updateAtData.ID] = weekdayMap[key]
	}

	// 更新星期几
	for i, v := range home.Modules[0].ContentList {
		v.Weekday = weekMap[v.ID]
		home.Modules[0].ContentList[i] = v
	}

	return home, nil
}

func (s *VideoService) Play(req req.PlayReq) (resp.PlayResp, error) {
	var err error
	var result resp.PlayResp

	// 构建多字段 term 查询
	termQuery1 := elastic.NewTermQuery("id", req.VideoID)
	termQuery2 := elastic.NewTermQuery("sid", req.PlayLine)
	termQuery3 := elastic.NewTermQuery("nid", req.EpisodeID)

	// 创建 Bool 查询
	boolQuery := elastic.NewBoolQuery()
	boolQuery = boolQuery.Must(termQuery1, termQuery2, termQuery3)

	resp, err := glob.Es.Search().Index("plays").Query(boolQuery).Do(context.Background())
	if err != nil {
		return result, err
	}

	if len(resp.Hits.Hits) == 0 {
		return result, nil
	}

	err = json.Unmarshal(resp.Hits.Hits[0].Source, &result)
	if err != nil {
		return result, err
	}

	numbers := extractNumber(result.LinkNext)
	if len(numbers) == 3 {
		id := strconv.Itoa(numbers[0])
		sid := strconv.Itoa(numbers[1])
		nid := strconv.Itoa(numbers[2])
		result.LinkNext = fmt.Sprintf("/play/%s/%s/%s", id, sid, nid)
	}

	return result, nil
}

// extractNumber 提取字符串中的数字
func extractNumber(input string) []int {
	// 定义一个匹配数字的正则表达式
	re := regexp.MustCompile(`\d+`)

	// 查找匹配的部分
	matches := re.FindAllString(input, -1)

	// 如果找到匹配的数字，则返回第一个匹配的结果
	if len(matches) > 0 {
		num1, _ := strconv.Atoi(matches[0])
		num2, _ := strconv.Atoi(matches[1])
		num3, _ := strconv.Atoi(matches[2])
		return []int{num1, num2, num3}
	}

	// 如果未找到匹配的数字，返回空字符串或其他合适的默认值
	return nil
}

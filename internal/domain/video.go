package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

type Video struct {
	model.Model
}

type PageVideoSearch struct {
	Video
	request.PageSearch
}

func (Video) TableName() string {
	return "video"
}
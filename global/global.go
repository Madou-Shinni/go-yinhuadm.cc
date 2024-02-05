package global

import (
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
	Es     *elastic.Client
)

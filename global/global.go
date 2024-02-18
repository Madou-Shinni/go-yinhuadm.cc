package global

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

var (
	Logger  *zap.Logger
	Es      *elastic.Client
	Redsync *redsync.Redsync
)

package global

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redsync/redsync/v4"
	"go.uber.org/zap"
)

var (
	Logger  *zap.Logger
	Es      *elasticsearch.TypedClient
	Redsync *redsync.Redsync
)

package common

import (
	"context"
	"errors"
	"github.com/goccy/go-json"
	"github.com/olivere/elastic/v7"
)

// MatchQuery 根据关键字搜索
func MatchQuery[T any](es *elastic.Client, index string, from, size int, searchFields []string, keyword string) (list []T, total int64, err error) {
	// 创建查询
	query := elastic.NewMultiMatchQuery(keyword, searchFields...)

	searchResult, err := es.Search().
		Index(index).
		Query(query).
		From(from).
		Size(size).
		Do(context.TODO())
	if err != nil {
		return
	}
	if searchResult.Hits == nil {
		err = errors.New("searchResult.Hits is nil")
		return
	}

	total = searchResult.TotalHits()

	for _, hit := range searchResult.Hits.Hits {
		var data T
		err = json.Unmarshal(hit.Source, &data)
		if err != nil {
			return
		}
		list = append(list, data)
	}

	return
}

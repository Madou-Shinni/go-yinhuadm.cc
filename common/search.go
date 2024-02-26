package common

import (
	"context"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/goccy/go-json"
	"strings"
)

// MatchQuery 根据关键字搜索
func MatchQuery[T any](es *elasticsearch.TypedClient, index string, from, size int, searchFields []string, keyword string, excludeFields []string, highlightFields []string) (list []T, total int64, err error) {

	hlFields := make(map[string]types.HighlightField, len(highlightFields))

	for _, field := range highlightFields {
		hlFields[field] = types.HighlightField{}
	}

	searchResult, err := es.Search().
		Index(index).
		Request(&search.Request{
			Query: &types.Query{
				MultiMatch: &types.MultiMatchQuery{
					Query:  keyword,
					Fields: searchFields,
				},
			},
			Highlight: &types.Highlight{
				Fields: hlFields,
			},
		}).
		SourceExcludes_(excludeFields...).
		From(from).
		Size(size).
		Do(context.TODO())
	if err != nil {
		return
	}
	if searchResult.Hits.Hits == nil {
		err = errors.New("searchResult.Hits is nil")
		return
	}

	total = searchResult.Hits.Total.Value

	for _, hit := range searchResult.Hits.Hits {
		var data T
		var dataMap map[string]interface{}
		var marshal []byte

		// 高亮
		if hit.Highlight != nil {
			err = json.Unmarshal(hit.Source_, &dataMap)
			if err != nil {
				return
			}

			for s, ss := range hit.Highlight {
				dataMap[s] = strings.Join(ss, "")
			}

			marshal, err = json.Marshal(dataMap)
			if err != nil {
				return
			}

			err = json.Unmarshal(marshal, &data)
			if err != nil {
				return
			}
		} else {
			err = json.Unmarshal(hit.Source_, &data)
			if err != nil {
				return
			}
		}

		list = append(list, data)
	}

	return
}

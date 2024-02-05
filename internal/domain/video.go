package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"time"
)

// Video 视频
type Video struct {
	ID           int64          `json:"id,omitempty"`
	Title        string         `json:"title,omitempty"`
	Tags         []string       `json:"tags,omitempty"`
	Cover        string         `json:"cover,omitempty"`
	Introduction string         `json:"introduction,omitempty"`
	Director     string         `json:"director,omitempty"`
	Screenwriter string         `json:"screenwriter,omitempty"`
	Note         string         `json:"note,omitempty"`
	Starrings    []string       `json:"starrings,omitempty"`
	Thirdlink    ThirdPartyLink `json:"thirdlink"`
	UpdateAt     time.Time      `json:"updateAt"`
	EpisodeList  []Episode      `json:"episodeList,omitempty"`
}

func (Video) Index() string {
	return "videos"
}

func (Video) Mappings() string {
	return `
{
	"mappings": {
      "properties": {
        "introduction": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "cover": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "director": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "episodeList": {
          "properties": {
            "episode": {
              "type": "long"
            },
            "playLine": {
              "type": "long"
            }
          }
        },
        "id": {
          "type": "long"
        },
        "note": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "screenwriter": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "starrings": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "tags": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "thirdlink": {
          "properties": {
            "link": {
              "type": "text",
              "fields": {
                "keyword": {
                  "type": "keyword",
                  "ignore_above": 256
                }
              }
            },
            "title": {
              "type": "text",
              "fields": {
                "keyword": {
                  "type": "keyword",
                  "ignore_above": 256
                }
              }
            }
          }
        },
        "title": {
          "type": "text",
		  "analyzer": "ik_smart"
        },
        "updateAt": {
          "type": "date"
        }
      }
    }
  }`
}

type PageVideoSearch struct {
	request.PageSearch
}

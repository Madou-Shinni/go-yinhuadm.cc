package domain

import "github.com/Madou-Shinni/gin-quickstart/constants"

// Episode 剧集
type Episode struct {
	ID       int64              `json:"ID,omitempty"` // 剧集这个id没有实际用处
	VideoID  int64              `json:"videoID,omitempty"`
	Episode  int                `json:"episode,omitempty"` // 剧集
	PlayLine constants.PlayLine `json:"playLine,omitempty"`
}

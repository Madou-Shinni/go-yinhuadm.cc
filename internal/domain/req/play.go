package req

type PlayReq struct {
	EpisodeID int `json:"episodeId,omitempty" form:"episodeId"`
	PlayLine  int `json:"playLine,omitempty" form:"playLine"`
	VideoID   int `json:"videoId,omitempty" form:"videoId"`
}

package domain

// Play 播放
type Play struct {
	Flag     string `json:"flag"`
	Encrypt  int    `json:"encrypt"`
	Trysee   int    `json:"trysee"`
	Points   int    `json:"points"`
	Link     string `json:"link"`
	LinkNext string `json:"link_next"`
	LinkPre  string `json:"link_pre"`
	VodData  struct {
		VodName     string `json:"vod_name"`
		VodActor    string `json:"vod_actor"`
		VodDirector string `json:"vod_director"`
		VodClass    string `json:"vod_class"`
	} `json:"vod_data"`
	URL     string `json:"url"`
	URLNext string `json:"url_next"`
	From    string `json:"from"`
	Server  string `json:"server"`
	Note    string `json:"note"`
	ID      string `json:"id"`  // videoID
	Sid     int    `json:"sid"` // playLine
	Nid     int    `json:"nid"` // episode
}

func (Play) Index() string {
	return "plays"
}

func (Play) Mappings() string {
	return ""
}

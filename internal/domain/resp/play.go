package resp

type PlayResp struct {
	Url      string `json:"url,omitempty"`
	UrlNext  string `json:"url_next,omitempty"`
	LinkNext string `json:"link_next,omitempty"`
}

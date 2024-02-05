package domain

// Module 模块
type Module struct {
	ID           int64     `json:"id,omitempty"`     // 模块ID
	Name         string    `json:"name,omitempty"`   // 模块名称
	Tabs         []string  `json:"tabs,omitempty"`   // 模块标签
	EnName       string    `json:"enName,omitempty"` // 模块英文名称
	ContentList  []Content `json:"contentList,omitempty"`
	ChildrenList []Module  `json:"childrenList,omitempty"`
}

// Content 内容
type Content struct {
	ID    int    `json:"id,omitempty"`
	Note  string `json:"note,omitempty"`
	Cover string `json:"cover,omitempty"`
	Title string `json:"title,omitempty"`
	Sort  int    `json:"sort,omitempty"`
}

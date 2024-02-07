package resp

import "github.com/Madou-Shinni/gin-quickstart/internal/domain"

type Home struct {
	Banners []domain.Banner `json:"banners,omitempty"`
	Modules []domain.Module `json:"modules,omitempty"`
}

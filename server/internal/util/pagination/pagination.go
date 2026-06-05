package pagination

const (
	DefaultLimit = 20
	MaxLimit     = 100
)

type Params struct {
	Limit  int `form:"limit" validate:"min=1,max=100"`
	Offset int `form:"offset" validate:"min=0"`
}

type Response[T any] struct {
	Data    []T   `json:"data"`
	Total   int64 `json:"total"`
	Limit   int   `json:"limit"`
	Offset  int   `json:"offset"`
	HasMore bool  `json:"has_more"`
}

func (p *Params) SetDefaults() {

	if p.Limit == 0 {
		p.Limit = DefaultLimit
	}

	if p.Limit > MaxLimit {
		p.Limit = MaxLimit
	}
}
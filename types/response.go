package types

type Response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Meta    Meta   `json:"meta,omitempty"`
	Type    string `json:"type,omitempty"`
	Code    string `json:"code,omitempty"`
}

type Meta struct {
	Total     int    `json:"total,omitempty"`
	Skipped   int    `json:"skipped,omitempty"`
	PerPage   int    `json:"perPage,omitempty"`
	Page      int    `json:"page,omitempty"`
	PageCount int    `json:"pageCount,omitempty"`
	Next      string `json:"next,omitempty"`
	Previous  string `json:"previous,omitempty"`
	NextStep  string `json:"nextStep,omitempty"`
}

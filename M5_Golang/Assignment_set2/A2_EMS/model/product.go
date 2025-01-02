package model

type Item struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Details     string  `json:"details"`
	Cost        float64 `json:"cost"`
	Quantity    int     `json:"quantity"`
	CategoryRef int     `json:"category_ref"`
}

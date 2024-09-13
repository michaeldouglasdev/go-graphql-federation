package model

type ItemReviewSummary struct {
	Total   int           `json:"total"`
	Average float64       `json:"average"`
	Review  []*ItemReview `json:"review"`
}

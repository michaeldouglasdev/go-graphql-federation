package model

type Item struct {
	ID            string             `json:"id"`
	ReviewSummary *ItemReviewSummary `json:"reviewSummary,omitempty"`
}

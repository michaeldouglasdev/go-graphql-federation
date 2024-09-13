package model

func (Item) IsEntity() {}

type ItemReview struct {
	User      *User   `json:"user"`
	Rating    int     `json:"rating"`
	CreatedAt string  `json:"createdAt"`
	Text      *string `json:"text,omitempty"`
}

func (ItemReview) IsCreateItemReviewResponse() {}

/*type CreateItemReviewResponse interface {
	IsCreateItemReviewResponse()
}

type CreateItemReviewError struct {
	Message string `json:"message"`
}*/

//func (CreateItemReviewError) IsCreateItemReviewResponse() {}

/*type CreateItemReviewInput struct {
	Rating int     `json:"rating"`
	Text   *string `json:"text,omitempty"`
}
*/

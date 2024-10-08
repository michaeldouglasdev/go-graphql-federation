// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateItemReviewResponse interface {
	IsCreateItemReviewResponse()
}

type CreateItemReviewError struct {
	Message string `json:"message"`
}

func (CreateItemReviewError) IsCreateItemReviewResponse() {}

type ItemReviewsMutations struct {
	Create CreateItemReviewResponse `json:"create"`
}

type LoginRequiredError struct {
	Message string `json:"message"`
}

func (LoginRequiredError) IsCreateItemReviewResponse() {}

type Mutation struct {
}

type Query struct {
}

type ReviewsQueries struct {
	List []*ItemReview `json:"list"`
}

type UnauthorizedError struct {
	Message string `json:"message"`
}

func (UnauthorizedError) IsCreateItemReviewResponse() {}

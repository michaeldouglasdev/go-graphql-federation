package model

type CartsMutations struct {
	Add    CartsAddResponse    `json:"add"`
	Update CartsUpdateResponse `json:"update"`
	Remove *CartsRemove        `json:"remove"`
	ID     string              `json:"id"` // to pass arg to child field
}

package model

type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Type     UserType `json:"type"`
}

func (User) IsEntity()              {}
func (User) IsGetUserResponse()     {}
func (User) IsMeUserResponse()      {}
func (User) IsUsersCreateResponse() {}

type UserNotFoundError struct {
	Message string `json:"message"`
}

func (UserNotFoundError) IsGetUserResponse() {}

/*type LoginUserInvalidUsernameOrPasswordError struct {
	Message string `json:"message"`
}

func (LoginUserInvalidUsernameOrPasswordError) IsGetUserResponse() {}*/

type GetUserResponse interface {
	IsGetUserResponse()
}

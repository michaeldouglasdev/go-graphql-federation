package model

import (
	"fmt"
	"io"
	"strconv"
)

type User struct {
	ID string `json:"id"`
}

func (User) IsEntity() {}

type UserType string

const (
	UserTypeBasic UserType = "BASIC"
	UserTypeAdmin UserType = "ADMIN"
)

func (e UserType) IsValid() bool {
	switch e {
	case UserTypeBasic, UserTypeAdmin:
		return true
	}
	return false
}

func (e UserType) String() string {
	return string(e)
}

func (e *UserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}
	return nil
}

func (e UserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

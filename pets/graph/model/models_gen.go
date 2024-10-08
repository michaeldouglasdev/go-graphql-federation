// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type PetsCreateResponse interface {
	IsPetsCreateResponse()
}

type PetsDeleteResponse interface {
	IsPetsDeleteResponse()
}

type PetsGetResponse interface {
	IsPetsGetResponse()
}

type Error struct {
	Message string `json:"message"`
}

func (Error) IsPetsGetResponse() {}

func (Error) IsPetsCreateResponse() {}

func (Error) IsPetsDeleteResponse() {}

type Mutation struct {
}

type Pet struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Birthday string   `json:"birthday"`
	Breed    PetBreed `json:"breed"`
	User     *User    `json:"user"`
}

func (Pet) IsPetsGetResponse() {}

func (Pet) IsPetsCreateResponse() {}

func (Pet) IsPetsDeleteResponse() {}

func (Pet) IsEntity() {}

type PetsCreateInput struct {
	Name     string   `json:"name"`
	Birthday string   `json:"birthday"`
	Breed    PetBreed `json:"breed"`
}

type PetsMutations struct {
	Create PetsCreateResponse `json:"create"`
	Delete PetsDeleteResponse `json:"delete,omitempty"`
}

type PetsQueries struct {
	All []*Pet          `json:"all"`
	Get PetsGetResponse `json:"get"`
}

type Query struct {
}

type User struct {
	ID string `json:"id"`
}

func (User) IsEntity() {}

type PetBreed string

const (
	PetBreedDog  PetBreed = "DOG"
	PetBreedCash PetBreed = "CASH"
)

var AllPetBreed = []PetBreed{
	PetBreedDog,
	PetBreedCash,
}

func (e PetBreed) IsValid() bool {
	switch e {
	case PetBreedDog, PetBreedCash:
		return true
	}
	return false
}

func (e PetBreed) String() string {
	return string(e)
}

func (e *PetBreed) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PetBreed(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PetBreed", str)
	}
	return nil
}

func (e PetBreed) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

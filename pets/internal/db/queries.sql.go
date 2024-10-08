// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
)

const createPet = `-- name: CreatePet :exec
INSERT INTO pets (id, name, birthday, breed, user_id)
VALUES (?, ?, ?, ?, ?)
`

type CreatePetParams struct {
	ID       string
	Name     string
	Birthday string
	Breed    PetsBreed
	UserID   string
}

func (q *Queries) CreatePet(ctx context.Context, arg CreatePetParams) error {
	_, err := q.db.ExecContext(ctx, createPet,
		arg.ID,
		arg.Name,
		arg.Birthday,
		arg.Breed,
		arg.UserID,
	)
	return err
}

const getPet = `-- name: GetPet :one
SELECT id, name, birthday, breed, user_id FROM pets WHERE ID = ?
`

func (q *Queries) GetPet(ctx context.Context, id string) (Pet, error) {
	row := q.db.QueryRowContext(ctx, getPet, id)
	var i Pet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Birthday,
		&i.Breed,
		&i.UserID,
	)
	return i, err
}

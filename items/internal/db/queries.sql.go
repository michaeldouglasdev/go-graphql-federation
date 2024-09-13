// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"
)

const createItem = `-- name: CreateItem :execresult
INSERT INTO items (
  id,
  name,
  price,
  image,
  expiration_time,
  weight,
  suitable_for,
  material,
  type
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateItemParams struct {
	ID             string
	Name           string
	Price          float64
	Image          string
	ExpirationTime sql.NullString
	Weight         sql.NullString
	SuitableFor    NullItemsSuitableFor
	Material       sql.NullString
	Type           ItemsType
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createItem,
		arg.ID,
		arg.Name,
		arg.Price,
		arg.Image,
		arg.ExpirationTime,
		arg.Weight,
		arg.SuitableFor,
		arg.Material,
		arg.Type,
	)
}

const getItem = `-- name: GetItem :one
SELECT id, name, price, image, expiration_time, weight, suitable_for, material, type FROM items WHERE id = ?
`

func (q *Queries) GetItem(ctx context.Context, id string) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Image,
		&i.ExpirationTime,
		&i.Weight,
		&i.SuitableFor,
		&i.Material,
		&i.Type,
	)
	return i, err
}

const listFoods = `-- name: ListFoods :many
SELECT id, name, price, image, expiration_time, weight, suitable_for, material, type FROM items WHERE type = "FOOD"
`

func (q *Queries) ListFoods(ctx context.Context) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listFoods)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Image,
			&i.ExpirationTime,
			&i.Weight,
			&i.SuitableFor,
			&i.Material,
			&i.Type,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listItems = `-- name: ListItems :many
SELECT id, name, price, image, expiration_time, weight, suitable_for, material, type FROM items
`

func (q *Queries) ListItems(ctx context.Context) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Image,
			&i.ExpirationTime,
			&i.Weight,
			&i.SuitableFor,
			&i.Material,
			&i.Type,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listToys = `-- name: ListToys :many
SELECT id, name, price, image, expiration_time, weight, suitable_for, material, type FROM items WHERE type = "TOY"
`

func (q *Queries) ListToys(ctx context.Context) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listToys)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Image,
			&i.ExpirationTime,
			&i.Weight,
			&i.SuitableFor,
			&i.Material,
			&i.Type,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

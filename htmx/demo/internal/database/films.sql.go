// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: films.sql

package database

import (
	"context"
)

const createFilm = `-- name: CreateFilm :one
INSERT INTO films (title, director)
VALUES ($1, $2)
RETURNING id, created_at, updated_at, title, director
`

type CreateFilmParams struct {
	Title    string `json:"title"`
	Director string `json:"director"`
}

func (q *Queries) CreateFilm(ctx context.Context, arg CreateFilmParams) (Film, error) {
	row := q.db.QueryRowContext(ctx, createFilm, arg.Title, arg.Director)
	var i Film
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Director,
	)
	return i, err
}

const getFilms = `-- name: GetFilms :many
SELECT id, created_at, updated_at, title, director FROM films
`

func (q *Queries) GetFilms(ctx context.Context) ([]Film, error) {
	rows, err := q.db.QueryContext(ctx, getFilms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Film
	for rows.Next() {
		var i Film
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Director,
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

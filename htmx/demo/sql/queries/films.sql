-- name: CreateFilm :one
INSERT INTO films (title, director)
VALUES ($1, $2)
RETURNING *;

-- name: GetFilms :many
SELECT * FROM films;

-- +goose Up
CREATE TABLE films (
  id UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

  title TEXT NOT NULL,
  director TEXT NOT NULL
);

-- +goose Down
DROP TABLE films;

-- +goose Up
CREATE TABLE Users(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL,
  api_key TEXT NOT NULL
);

-- +goose Down
DROP TABLE Users;

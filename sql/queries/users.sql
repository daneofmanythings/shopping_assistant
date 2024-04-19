-- name: UserCreate :one
INSERT INTO Users ( id, created_at, updated_at, name, api_key ) 
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: UserGet :one
SELECT * FROM Users
WHERE api_key is ?;

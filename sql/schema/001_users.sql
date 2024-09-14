-- +goose Up
CREATE TABLE users (
	id uuid primary key,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	name VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE users;

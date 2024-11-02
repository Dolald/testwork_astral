-- +goose Up
-- +goose StatementBegin

CREATE TABLE users
(
    id serial not null unique,
    login varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE documents
(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    filename VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE documents;
-- +goose StatementEnd




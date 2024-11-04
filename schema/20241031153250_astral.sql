-- +goose Up
-- +goose StatementBegin

CREATE TABLE users
(
    id SERIAL NOT NULL,
    login VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
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




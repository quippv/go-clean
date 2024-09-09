-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    created_at BIGINT NOT NULL DEFAULT (extract(epoch from now()) * 1000)::BIGINT,
    updated_at BIGINT NOT NULL DEFAULT (extract(epoch from now()) * 1000)::BIGINT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

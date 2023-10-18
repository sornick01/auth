-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE users(
    id serial PRIMARY KEY,
    name VARCHAR,
    email VARCHAR,
    password VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE users;
-- +goose StatementEnd


-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users(
    user_id SERIAL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    name varchar(20),
    email varchar(20)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE users;
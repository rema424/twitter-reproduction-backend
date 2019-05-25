
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE users ADD user_name VARCHAR(20) UNIQUE AFTER user_id;

ALTER TABLE users ADD (
    icon VARCHAR(255),
    header_image VARCHAR(255),
    profile VARCHAR(150),
    birthday DATE,
    place VARCHAR(100),
    url VARCHAR(255)
);

UPDATE users SET user_name = 'alice' WHERE name = 'Alice' AND email = 'alice@example.com';
UPDATE users SET user_name = 'bob' WHERE name = 'Bob' AND email = 'bob@example.com';
UPDATE users SET user_name = 'carol' WHERE name = 'Carol' AND email = 'carol@example.com';
UPDATE users SET user_name = 'dave' WHERE name = 'Dave' AND email = 'dave@example.com';
UPDATE users SET user_name = 'eve' WHERE name = 'Eve' AND email = 'eve@example.com';
UPDATE users SET user_name = 'frank' WHERE name = 'Frank' AND email = 'frank@example.com';
UPDATE users SET user_name = 'grace' WHERE name = 'Grace' AND email = 'grace@example.com';
UPDATE users SET user_name = 'heidi' WHERE name = 'Heidi' AND email = 'heidi@example.com';
UPDATE users SET user_name = 'ivan' WHERE name = 'Ivan' AND email = 'ivan@example.com';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE users
DROP user_name,
DROP icon,
DROP header_image,
DROP profile,
DROP birthday,
DROP place,
DROP url;

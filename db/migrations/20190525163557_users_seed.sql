
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO users (name, email) VALUES
('Alice', 'alice@example.com'),
('Bob', 'bob@example.com'),
('Carol', 'carol@example.com'),
('Dave', 'dave@example.com'),
('Eve', 'eve@example.com'),
('Frank', 'frank@example.com'),
('Grace', 'grace@example.com'),
('Heidi', 'heidi@example.com'),
('Ivan', 'ivan@example.com');


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM users
WHERE (name = 'Alice' AND email = 'alice@example.com')
OR (name = 'Bob' AND email = 'bob@example.com')
OR (name = 'Carol' AND email = 'carol@example.com')
OR (name = 'Dave' AND email = 'dave@example.com')
OR (name = 'Eve' AND email = 'eve@example.com')
OR (name = 'Frank' AND email = 'frank@example.com')
OR (name = 'Grace' AND email = 'grace@example.com')
OR (name = 'Heidi' AND email = 'heidi@example.com')
OR (name = 'Ivan' AND email = 'ivan@example.com');
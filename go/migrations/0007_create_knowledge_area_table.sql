-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE knowledge_area
(
    cipher VARCHAR NOT NULL UNIQUE,
    title  VARCHAR NOT NULL,
    PRIMARY KEY (cipher, title)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS knowledge_area;

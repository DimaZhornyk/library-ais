-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE book_area
(
    book_isbn   VARCHAR NOT NULL CONSTRAINT book_isbn_fk REFERENCES books ON DELETE CASCADE,
    area_cipher VARCHAR NOT NULL CONSTRAINT area_cipher_fk REFERENCES knowledge_area (cipher) ON DELETE CASCADE
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS book_area;

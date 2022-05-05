-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE book_authors
(
    id        SERIAL CONSTRAINT book_author_pk PRIMARY KEY,
    surname   VARCHAR     NOT NULL,
    book_isbn VARCHAR NOT NULL CONSTRAINT book_isbn_fk REFERENCES books
);

CREATE UNIQUE INDEX book_author_id_uindex ON book_authors (id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS book_authors;
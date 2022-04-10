-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE books
(
    isbn             VARCHAR          NOT NULL CONSTRAINT books_pk PRIMARY KEY,
    title            VARCHAR          NOT NULL,
    city             VARCHAR          NOT NULL,
    publishing_house VARCHAR          NOT NULL,
    year             SMALLINT         NOT NULL,
    pages_quantity   INTEGER          NOT NULL,
    price            DOUBLE PRECISION NOT NULL
);

CREATE UNIQUE INDEX books_isbn_uindex ON books (isbn);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS books;
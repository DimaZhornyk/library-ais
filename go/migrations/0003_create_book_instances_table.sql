-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE book_instances
(
    inventory_number VARCHAR NOT NULL CONSTRAINT book_instances_pk PRIMARY KEY,
    book_isbn        VARCHAR NOT NULL CONSTRAINT book_isbn_fk REFERENCES books ON DELETE CASCADE,
    shelf            VARCHAR NOT NULL
);

CREATE UNIQUE INDEX book_instances_inventory_number_uindex ON book_instances (inventory_number);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS book_instances;
-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE checkouts
(
    checkout_number       SERIAL  CONSTRAINT checkouts_pk PRIMARY KEY,
    reader_card_number    VARCHAR NOT NULL,
    book_inventory_number VARCHAR NOT NULL CONSTRAINT book_inventory_number_fk  REFERENCES book_instances ON UPDATE CASCADE,
    checkout_date         DATE    NOT NULL,
    expected_return_date  DATE    NOT NULL,
    return_date           DATE,
    repaid                DOUBLE PRECISION
);

CREATE UNIQUE INDEX checkouts_checkout_number_uindex ON checkouts (checkout_number);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS checkouts;
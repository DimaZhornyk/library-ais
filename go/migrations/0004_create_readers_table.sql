-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE readers
(
    card_number  VARCHAR NOT NULL CONSTRAINT reader_pk PRIMARY KEY,
    full_name    VARCHAR NOT NULL,
    home_address VARCHAR NOT NULL,
    seat         VARCHAR NOT NULL,
    birth_date   DATE    NOT NULL check ( birth_date < (current_date - interval '17' year ) )
);

CREATE UNIQUE INDEX readers_card_number_uindex ON readers (card_number);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS readers;
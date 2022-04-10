-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE readers_phones
(
    id          SERIAL CONSTRAINT readers_phones_pk PRIMARY KEY,
    reader_card VARCHAR NOT NULL CONSTRAINT reader_card_fk REFERENCES readers ON DELETE CASCADE,
    phone       VARCHAR NOT NULL
);

CREATE UNIQUE INDEX readers_phones_id_uindex ON readers_phones (id);

CREATE UNIQUE INDEX readers_phones_phone_uindex ON readers_phones (phone);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS readers_phones;
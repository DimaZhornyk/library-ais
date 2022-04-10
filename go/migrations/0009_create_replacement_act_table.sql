-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE replacement_act
(
    id                   SERIAL CONSTRAINT replacement_act_pk PRIMARY KEY,
    replacement_date     DATE    NOT NULL,
    old_inventory_number VARCHAR NOT NULL CONSTRAINT old_inventory_number_fk REFERENCES book_instances,
    new_inventory_number VARCHAR NOT NULL CONSTRAINT new_inventory_number_fk REFERENCES book_instances
);

CREATE UNIQUE INDEX replacement_act_id_uindex ON replacement_act (id);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS replacement_act;
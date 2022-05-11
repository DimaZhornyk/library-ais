-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose StatementBegin

CREATE OR REPLACE FUNCTION delete_reader(_card_number VARCHAR) RETURNS BOOLEAN AS
$$
BEGIN
    IF EXISTS(
            SELECT *
            FROM checkouts
            WHERE return_date IS NULL
              AND reader_card_number = _card_number
        ) THEN
        RAISE EXCEPTION 'The debt exists';
    ELSE
        DELETE
        FROM readers
        WHERE card_number = _card_number;
    END IF;

    RETURN true;
END;
$$
LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose StatementBegin
create or replace function is_valid_checkout(
    _reader_card_number varchar,
    _book_inventory_number varchar
) RETURNS BOOLEAN AS
$$
DECLARE book_with_in int;
    DECLARE cnt_same_books_user int;
    DECLARE cnt_checkouts_of_book int;
    DECLARE _book_isbn varchar;
BEGIN
    SELECT MIN(book_isbn), COUNT(*) INTO _book_isbn, book_with_in
    FROM book_instances
    WHERE inventory_number = _book_inventory_number;

    -- book with this inventory number was taken
    SELECT COUNT(*) INTO cnt_checkouts_of_book
    FROM checkouts INNER JOIN book_instances bi ON checkouts.book_inventory_number = bi.inventory_number
    WHERE book_inventory_number = _book_inventory_number AND return_date IS NULL;

    -- user already took the book with the same isbn
    SELECT COUNT(*) INTO cnt_same_books_user
    FROM checkouts INNER JOIN book_instances bi2 ON bi2.inventory_number = checkouts.book_inventory_number
    WHERE reader_card_number = _reader_card_number AND return_date IS NULL AND book_isbn = _book_isbn;

    RETURN book_with_in > 0 AND cnt_same_books_user = 0 AND cnt_checkouts_of_book = 0;
END;
$$
language plpgsql;

SELECT create_constraint_if_not_exists('checkouts', 'create_checkout_constraint', ' CHECK(is_valid_checkout(reader_card_number, book_inventory_number));');

-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

ALTER TABLE checkouts DROP CONSTRAINT create_checkout_constraint;
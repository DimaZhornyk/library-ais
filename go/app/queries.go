package app

var queries = []string{
	// create books table
	`CREATE TABLE books
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
	`,

	// create book_authors table
	`CREATE TABLE book_authors
(
	id        SERIAL CONSTRAINT book_author_pk PRIMARY KEY,
	surname   INT     NOT NULL,
	book_isbn VARCHAR NOT NULL CONSTRAINT book_isbn_fk REFERENCES books
);

	CREATE UNIQUE INDEX book_author_id_uindex ON book_author (id);
`,

	// create book_instances table
	`CREATE TABLE book_instances
(
    inventory_number VARCHAR NOT NULL CONSTRAINT book_instances_pk PRIMARY KEY,
    book_isbn        VARCHAR NOT NULL CONSTRAINT book_isbn_fk REFERENCES books ON DELETE CASCADE,
    shelf            VARCHAR NOT NULL
);

CREATE UNIQUE INDEX book_instances_inventory_number_uindex ON book_instances (inventory_number);
`,

	// create readers table
	`CREATE TABLE readers
(
    card_number  VARCHAR NOT NULL CONSTRAINT reader_pk PRIMARY KEY,
    full_name    VARCHAR NOT NULL,
    home_address VARCHAR NOT NULL,
    seat         VARCHAR NOT NULL,
    birth_date   DATE    NOT NULL check ( birth_date < (current_date - interval '17' year ) )
);

CREATE UNIQUE INDEX reader_card_number_uindex ON reader (card_number);
`,

	// create readers_phones table
	`CREATE TABLE readers_phones
(
    id          SERIAL CONSTRAINT readers_phones_pk PRIMARY KEY,
    reader_card VARCHAR NOT NULL CONSTRAINT reader_card_fk REFERENCES readers ON DELETE CASCADE,
    phone       VARCHAR NOT NULL
);

CREATE UNIQUE INDEX readers_phones_id_uindex ON readers_phones (id);

CREATE UNIQUE INDEX readers_phones_phone_uindex ON readers_phones (phone);
`,

	// create checkouts table
	`CREATE TABLE checkouts
(
    checkout_number      INTEGER NOT NULL CONSTRAINT checkouts_pk PRIMARY KEY,
    checkout_date        DATE    NOT NULL,
    expected_return_date DATE    NOT NULL,
    return_date          DATE,
    repaid               DOUBLE PRECISION
);

CREATE UNIQUE INDEX checkouts_checkout_number_uindex ON checkouts (checkout_number);
`,
}

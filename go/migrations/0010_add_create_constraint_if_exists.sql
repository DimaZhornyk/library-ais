-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose StatementBegin

CREATE OR REPLACE FUNCTION create_constraint_if_not_exists (t_name text, c_name text, constraint_sql text)
    RETURNS void
AS
$BODY$
begin
    -- Look for our constraint
    if not exists (select constraint_name
                   from information_schema.constraint_column_usage
                   where table_name = t_name  and constraint_name = c_name) then
        execute 'ALTER TABLE ' || t_name || ' ADD CONSTRAINT ' || c_name || ' ' || constraint_sql;
    end if;
end;
$BODY$
    LANGUAGE plpgsql VOLATILE;

-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
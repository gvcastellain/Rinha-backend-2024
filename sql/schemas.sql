CREATE TYPE transaction_type as ENUM ('c', 'd')

CREATE TABLE IF NOT EXISTS CLIENT (
    id SERIAL PRIMARY KEY NOT NULL,
    current_balance integer not NULL,
    client_limit integer not NULL
)

CREATE TABLE IF NOT EXISTS TRANSACTION (
    id SERIAL PRIMARY KEY NOT NULL,
    client_id integer NOT NULL,
    description varchar(10),
    transaction_date timestamp,
    value integer,
    CONSTRAINT fk_client_transaction FOREIGN KEY client_id REFERENCES CLIENT(id) ON DELETE CASCADE
)
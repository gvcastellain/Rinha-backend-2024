CREATE TYPE transaction_type as ENUM ('c', 'd')

CREATE TABLE IF NOT EXISTS client (
    id SERIAL PRIMARY KEY NOT NULL,
    current_balance integer not NULL,
    client_limit integer not NULL
)

CREATE TABLE IF NOT EXISTS transaction (
    id SERIAL PRIMARY KEY NOT NULL,
    client_id integer NOT NULL,
    description varchar(10),
    transaction_date timestamp,
    value integer,
    transaction_type transaction_type
    CONSTRAINT fk_client_transaction FOREIGN KEY client_id REFERENCES client(id) ON DELETE CASCADE
)

INSERT INTO client (current_balance, client_limit)
VALUES 
(0, 100000),
(0, 80000),
(0, 1000000),
(0, 10000000),
(0, 500000);    
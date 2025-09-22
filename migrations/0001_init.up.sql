CREATE TABLE IF NOT EXISTS users (
    id varchar(36) NOT NULL PRIMARY KEY,
    name text NOT NULL,
    email text NOT NULL,
    password varchar(60) NOT NULL,
    UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS tasks (
    id varchar(36) NOT NULL PRIMARY KEY,
    name text NOT NULL,
    description text NOT NULL,
    status text NOT NULL
);
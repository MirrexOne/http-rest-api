CREATE TABLE users
(
    id                 BIGSERIAL    NOT NULL PRIMARY KEY,
    email              VARCHAR(255) NOT NULL UNIQUE,
    encrypted_password VARCHAR      NOT NULL
);
CREATE TABLE url (
                     id CHAR(6) PRIMARY KEY,
                     url VARCHAR NOT NULL,
                     email VARCHAR(255) NOT NULL,
                     visit INTEGER NOT NULL DEFAULT 0,
                     is_del BOOLEAN NOT NULL DEFAULT FALSE,
                     UNIQUE(url)
);
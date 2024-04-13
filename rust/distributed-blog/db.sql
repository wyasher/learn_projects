CREATE TABLE categories
(
    id     SERIAL PRIMARY KEY,
    name   VARCHAR(100) NOT NULL,
    is_del BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE TABLE topics
(
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255)             NOT NULL,
    category_id INT                      NOT NULL,
    summary     VARCHAR(255)             NOT NULL,
    content     VARCHAR                  NOT NULL,
    hit         INT                      NOT NULL DEFAULT 0,
    dateline    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_del      BOOLEAN                  NOT NULL DEFAULT FALSE
);

CREATE TABLE admins
(
    id       SERIAL PRIMARY KEY,
    email    VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_del   BOOLEAN      NOT NULL DEFAULT FALSE
);
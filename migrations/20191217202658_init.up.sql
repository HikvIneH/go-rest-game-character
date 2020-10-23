CREATE TABLE album
(
    id         VARCHAR PRIMARY KEY,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE character
(
    id                      VARCHAR PRIMARY KEY,
    name                    VARCHAR NOT NULL,
    character_code          bigint NOT NULL,
    character_power         VARCHAR NOT NULL,
    character_value         VARCHAR NOT NULL,
    created_at              TIMESTAMP NOT NULL,
    updated_at              TIMESTAMP NOT NULL
);

-- CREATE TABLE character_value
-- (
--     id                  VARCHAR PRIMARY KEY,
--     character_id        VARCHAR ,
--     character_value     VARCHAR , 
--     created_at          TIMESTAMP NOT NULL,
--     updated_at          TIMESTAMP NOT NULL
-- );
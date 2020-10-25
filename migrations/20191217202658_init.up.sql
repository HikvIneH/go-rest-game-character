CREATE TABLE character_type
(
    character_code          bigint PRIMARY KEY,
    name                    VARCHAR NOT NULL,
    created_at              TIMESTAMP NOT NULL,
    updated_at              TIMESTAMP NOT NULL
);

CREATE TABLE character
(
    id                      VARCHAR PRIMARY KEY,
    name                    VARCHAR NOT NULL,
    character_code          bigint NOT NULL,
    character_power         bigint NOT NULL,
    character_value         bigint NOT NULL,
    created_at              TIMESTAMP NOT NULL,
    updated_at              TIMESTAMP NOT NULL,

    CONSTRAINT fk_customer
        FOREIGN KEY(character_code) 
        REFERENCES character_type(character_code)
);
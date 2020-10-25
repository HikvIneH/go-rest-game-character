CREATE TABLE character
(
    id                      VARCHAR PRIMARY KEY,
    name                    VARCHAR NOT NULL,
    character_code          bigint NOT NULL,
    character_power         bigint NOT NULL,
    character_value         bigint NOT NULL,
    created_at              TIMESTAMP NOT NULL,
    updated_at              TIMESTAMP NOT NULL

    CONSTRAINT fk_customer
        FOREIGN KEY(character_code) 
        REFERENCES charater_type(character_code)
);

CREATE TABLE character_type
(
    id                      bigint PRIMARY KEY,
    character_code          bigint NOT NULL,
    name                    VARCHAR NOT NULL,
    created_at              TIMESTAMP NOT NULL,
    updated_at              TIMESTAMP NOT NULL
)
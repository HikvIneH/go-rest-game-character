## Getting Started

This api is using starter kit from [qianxue go rest api](https://github.com/qiangxue/go-rest-ap)


[Docker](https://www.docker.com/get-started) is needed if you want to try the kit without setting up your
own database server. The kit requires **Docker 17.05 or higher** for the multi-stage build support.

[Docker Compose](https://docs.docker.com/compose/) is also needed to run docker-compose up -d

After installing Go and Docker, run the following commands to start experiencing this starter kit:

```shell
# start it with docker compose
docker-compose up -d

# to check on unit test 
~ make db-start
~ make test

```


## CRUD Operation
If you have `cURL` or some API client tools (e.g. [Postman](https://www.getpostman.com/)), you may try the following 
more complex scenarios:

```shell
# authenticate the user via: POST /v1/login
curl -X POST -H "Content-Type: application/json" -d '{"username": "demo", "password": "pass"}' http://localhost:8000/v1/login
# should return a JWT token like: {"token":"...JWT token here..."}

# with the above JWT token, access the album resources, such as: GET /v1/characters
curl -X GET -H "Authorization: Bearer ...JWT token here..." http://localhost:8000/v1/characters
# should return a list of album records in the JSON format

# Create character
curl -X POST -H "Authorization: Bearer ...JWT token here..." -H "Content-Type: application/json" -d '{"name":"Gandalf", "character_code":1, "character_power":100}' http://localhost:8000/v1/characters
# should return created character

# Update character
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Gandalf Update 10", "character_power":10}' http://localhost:8000/v1/characters/{ ID }
# should return updated character with ID
# should update character value based on newly updated characer power
```

## Database Schema

```
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
```
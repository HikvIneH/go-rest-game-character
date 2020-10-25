
# Go Rest API Game Character

[![Code Coverage](https://codecov.io/gh/HikvIneH/go-rest-game-character/branch/master/graph/badge.svg)](https://codecov.io/gh/HikvIneH/go-rest-game-character)

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
~ make test

```
Notes:
- Postgresql will use port 5432
- 


## CRUD Operation

RESTful API server running at `http://127.0.0.1:8000`. It provides the following endpoints:

* `GET /healthcheck`: a healthcheck service provided for health checking purpose (needed when implementing a server cluster)
* `POST /v1/login`: authenticates a user and generates a JWT
* `GET /v1/characters`: returns a paginated list of the characters
* `GET /v1/characters/:id`: returns the detailed information of an character
* `POST /v1/characters`: creates a new character
* `PUT /v1/characters/:id`: updates an existing character
* `DELETE /v1/characters/:id`: deletes an character


If you have `cURL` or some API client tools (e.g. [Postman](https://www.getpostman.com/)), you may try the following 
more complex scenarios:

```shell
# authenticate the user via: POST /v1/login
curl -X POST -H "Content-Type: application/json" -d '{"username": "demo", "password": "pass"}' http://localhost:8000/v1/login
# should return a JWT token like: {"token":"...JWT token here..."}

# with the above JWT token, access the character resources, such as: GET /v1/characters
curl -X GET -H "Authorization: Bearer ...JWT token here..." http://localhost:8000/v1/characters
# should return a list of character records in the JSON format

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
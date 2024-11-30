# simple dashboard app

> Note: This project aims to complete the assignment from MassiveMusic.co.id. 

simple dasboard in simple app that contains auth feature (register, login, forgot password, notif email, etc) and dummy dashboard page with edit user feature.

## How to Run

Here the tutorial to run the application

### Local

Prerequisite:
- install go with version 1.23 or up
- makefile runner
- postgresql database
- node JS with version 1.22 or up
- npm / yarn
- swagger cli
- mockery cli


Step for running server / service:
- change to directory `server`
- setup env with create `.env` file, refer to file `.env.example`
- run `go mod tidy`
- run `go run migration/migration.go` for migration table
- run `go run main.go`
- server must be running in local with port `:8080`

Step for running client:
- change to directory `client`
- setup env with create `.env` file, refer to file `.env.example`
- run `yarn install`
- run `yarn dev`
- client must be running in local with port `:3000`

Congrats, you can use the simple dashboard app

### Docker

Prerequisite:
- install docker
- install docker-compose

Step for running docker:
- go to file `docker-compose.yaml`, set the environment in `server-api`, based on `/server/.env.example`
- set the environment in `client-app`, based on `/client/.env.example`
- run `docker-compose up -d`


Congrats, you can use the simple dashboard app

## Unit testing

generate mockery for unit test server
```
mockery --all --dir "./repository" --output "./mocks/repository" --keeptree
mockery --all --dir "./service" --output "./mocks/service" --keeptree
```

Unit testing in server is focusing in repository and service layer

To running unit test in server, go to `server` directory and run `go test ./... -v`

To running unit testing in client, go to `client` directory and run `yarn test`

## github action

Im try to implement github action for pipelining or implement simple CI/CD.

For example to make sure all the unit test is `PASSED` or `OK`. Then checking for writing code, make sure the code is clean / effective.

Github action will be run when push new code in branch `main` or merge pull request to branch `main` with process:
- testing server
- lint server
- testing client
- lint client

## Tech stack

- golang for server / service main language
- typescript for client main languange
- Next js for frontend framework
- swagger for generate documentation api
- mockgen / mockery to generate mock code for golang testing
- docker for containerizing app
- docker compose for composing container app
- github action for pipelining
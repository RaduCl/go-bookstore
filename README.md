# Bookstore CRUD api

A basic bookstore CRUD api using golang, mysql(Mariadb)

## Requirement

docker, docker-compose, golang > 1.6

## Installation

Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
```bash
go mod download
```

set and load the .env variables onto the local host machine env
```bash
eval $(cat .env | sed 's/^/export /')
```

start the app
```bash
go run main.go
```

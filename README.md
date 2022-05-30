# Bookstore CRUD api

A basic bookstore CRUD api using golang, mysql(Mariadb)

## Requirement

docker, docker-compose, golang > 1.6

## Installation

set and load the .env variables onto the local host machine env
```bash
eval $(cat .env | sed 's/^/export /')
```

start the app
```bash
go run main.go
```

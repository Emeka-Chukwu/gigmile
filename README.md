# gigmile




## Architecture

This project using `Clean Architecture` with 4 domain layers:

- Model
- Data (repository)
- Handler

## How to run this project

## System Requirements

- Golang
- Docker
- Postgres (included in docker compose)

## Running

Setting up all containers

```console
$ make up_build
```

or 


```console
$ make up
```


Destroy running containers

```console
$ make down
```


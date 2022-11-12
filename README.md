# gigmile




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

## Endpoints

### Create a record (Post Method)
```console
$ /countries
```
### Get all records (Get Method)
```console
$ /countries
```
### Get a record (Get Method)
```console
$ /countries/{id}
```
### Update a record (Patch Method)
```console
$ /countries/{id}
```
### Delete a record (Delete Method)
```console
$ /countries/{id}
```



# johto API server

Simple customer project management written in Go.

A current version of [Go](https://golang.org/) and
[docker](https://www.docker.com/) is required. Start database with:

``` shell
docker-compose up -d db
```

Start HTTP server with:

``` bash
go run main.go
```

Use [postman](https://learning.getpostman.com/) for requests. A collection can
be found in `/docs/johto.postman_collection.json`.

# Go clean architecture sample

## Using

```sh
# db setup
docker run \
    --name my-postgres \
    -it --rm -p 5432:5432 \
    -e POSTGRES_PASSWORD=my-postgres \
    -e POSTGRES_USER=my-postgres \
    -e POSTGRES_DB=my-postgres \
    -d postgres

# server start
go run src/app/server.go

# healthCheck
curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST localhost:1323/healthCheck

# app endpoint
curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"title":"my title","body":"my body"}' localhost:1323/posts
curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST localhost:1323/posts/:id
curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST localhost:1323/posts
```

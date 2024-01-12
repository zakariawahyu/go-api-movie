# Api Movie ✨

### Technologies - Libraries
* [PostgreSQL](https://www.postgresql.org) - PostgreSQL Database
* [Redis](https://redis.com) - Redis
* [labstack/echo](https://go.dev/dl/) - Web Framework
* [spf13/viper](https://github.com/spf13/viper) - Config Reader
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Go Struct and Field validation
* [testify](https://github.com/stretchr/testify) - A toolkit with common assertions
* [sqlmock](https://github.com/DATA-DOG/go-sqlmock) - Mocking SQL Database

This app required 2 database connection
```
# run postgreSQL
docker run -d -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=api_movie postgres

# run redis
docker run -d -p 6379:6379 redis --requirepass "masukredis"
``` 

### Test
Run below command to run test, and make sure that all tests are passing
```
go test -v ./...
```

### Runing App
1. Copy config.example.yaml to config.yaml and fill the config
```azure
AppVersion:

Server:
  Port: :
  Development:
  ReadTimeout:
  WriteTimeout:

Postgres:
  Host:
  Port:
  User:
  Password:
  DbName:

Redis:
  Host:
  Port:
  Password:
```
2. Runing app docker-compose.yaml
```
make docker-compose
```
or
```
$ docker docker-compose up -d --build
```
- After run this, don't forget to check all container are successfully started


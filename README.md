# Go-Aut-Registration-OTP

## Short description
```
...
```

## Dependencies

go mod tidy
go mod vendor
## Initialization
Describe make command from Makefile
```
...
```

## Config

```
./config.json
```

You can also specify your own config file by run app with flag `-config`

## Migrations
installation 
go install github.com/pressly/goose/v3/cmd/goose@latest

Migration creation:
1) Choose migration directory
2) Create migration:
```shell script
$ goose create name_of_migration sql
```
3) In created file write your migration
4) Apply migration:
```shell script
 $ goose postgres "postgres://rpguser:rpguser@10.10.15.90:5452/rpgDB?sslmode=disable" up
# RCD Gopher Social

This is a Go app that simulates a social network. This is for a backend engineering with Go
course on Udemy.

[https://udemy.com/course/backend-engineering-with-go/](https://udemy.com/course/backend-engineering-with-go/)

## Migrations

To create a new migration, run the following command:

```bash
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users
```

To perform a migration, run the following command, replacing the database URL with the correct one:

```bash
migrate -path=./cmd/migrate/migrations/ -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" up
```

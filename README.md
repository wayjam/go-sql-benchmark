# Golang SQL-libs Benchmark

## Libs to test

This benchmark will use Postgres as a test database and [pgx](https://github.com/jackc/pgx) as the driver.

-	database/sql
-	sqlx
-	pgx
-	gorm
-	squirrel

## Run

**Start a PostgreSQL Image**:

```sh
docker run -d \
--name go-sql-benchmark-pg \
-p 5432:5432 \
-e POSTGRESQL_PASSWORD=postgres \
-e POSTGRESQL_DATABASE=test 
-e POSTGRESQL_USERNAME=postgres \
bitnami/postgresql:latest
```

**Start the test**

```sh
export BENCHMARK_SQL_DSN="host=localhost user=postgres password=postgres dbname=test sslmode=disable"
# run
go test  -v ./... -bench=. -benchmem
```

## Report

### Reference report

**Environment**

- MacBook Pro (15-inch, 2019) with 2.6 GHz 6-Core Intel Core i7 & 16GB RAM.
- PostgreSQL 9.6 in docker(darwin)

## TODO

Feel free to contribute to this project.

[ ] CI and auto-genereated report
[ ] prepare statement benchmark
[ ] [ent](https://github.com/facebook/ent) benchmark
[ ] code comments

# Golang SQL-libs Benchmark

## Libs to test

This benchmark will use Postgres as a test database and [pgx](https://github.com/jackc/pgx) as the driver.

-	database/sql
-	[sqlx](https://github.com/jmoiron/sqlx)
-	[pgx](https://github.com/jackc/pgx)
-	[gorm](https://github.com/go-gorm/gorm)
-	[squirrel](https://github.com/Masterminds/squirrel)

## Run

**Start a PostgreSQL Image**:

```sh
docker run -d \
--name go-sql-benchmark-pg \
-p 5432:5432 \
-e POSTGRES_PASSWORD=postgres \
-e POSTGRES_DB=test \
-e POSTGRES_USERNAME=postgres \
postgres:12
```

**Start the test**

```sh
export BENCHMARK_SQL_DSN="host=localhost user=postgres password=postgres dbname=test sslmode=disable"
# run
go test  -v ./... -bench=. -benchmem
```

## Report

Find the reference report at: <https://github.com/wayjam/go-sql-benchmark/actions?query=workflow%3A%22Go+SQL-libs+Benchmark%22>

## TODO

- [x] CI and auto-genereated report
- [ ] prepare statement benchmark
- [ ] [ent](https://github.com/facebook/ent) benchmark
- [ ] add more code comments

## Contribute

Feel free to contribute to this project.

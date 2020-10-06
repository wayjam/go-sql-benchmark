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

Using [cob](https://github.com/knqyf263/cob) to run benchmark on Github Actions.

Find the reference report at: <https://github.com/wayjam/go-sql-benchmark/actions?query=branch%3Amaster+is%3Asuccess>

## TODO

- [x] CI and auto-genereated report
- [ ] prepare statement benchmark
- [ ] [ent](https://github.com/facebook/ent) benchmark
- [ ] add more code comments
- [ ] comparison between libs.

## Contribute

Feel free to make contributiton to this project.

### Donate

<a href="https://www.buymeacoffee.com/wayjam" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Cola" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

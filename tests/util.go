package tests

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/wayjam/go-sql-benchmark/database"

	"github.com/jmoiron/sqlx"
	// _ "github.com/jackc/pgx/stdlib" // gorm imported, no need to import again
)

var (
	// benchmark options
	MAX_IDLE_CONN int
	MAX_CONN      int
	DSN           string

	// db for init and insert record
	db *sql.DB
)

func init() {
	DSN = os.Getenv("BENCHMARK_SQL_DSN")

	fmt.Println("Setup for whole benchmark...")

	var err error
	db, err = sql.Open("pgx", DSN)
	if err != nil {
		fmt.Printf("new db conn failed: %s", err.Error())
		os.Exit(2)
	}
	err = database.InitDB(db)
	if err != nil {
		fmt.Printf("Init db failed: %s", err.Error())
		os.Exit(2)
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func AssertBNoError(b *testing.B, err error) {
	if err != nil {
		b.Errorf("Unexpected error: %s", err.Error())
	}
}

func prepareWrapper(b *testing.B, f func()) {
	b.StopTimer()
	defer b.StartTimer()
	f()
}

func newSQLDB4Benchmark(b *testing.B) *sql.DB {
	db, err := sql.Open("pgx", DSN)
	AssertBNoError(b, err)

	err = db.Ping()
	AssertBNoError(b, err)

	return db
}

func newSQLXDB4Benchmark(b *testing.B) *sqlx.DB {
	db, err := sqlx.Open("pgx", DSN)
	AssertBNoError(b, err)

	err = db.Ping()
	AssertBNoError(b, err)

	return db
}

package tests

import (
	"testing"

	"github.com/wayjam/go-sql-benchmark/database"

	"github.com/jmoiron/sqlx"
)

var (
	sqlxDB *sqlx.DB
)

func sqlxSelectMulti(b *testing.B) {
	prepareWrapper(b, func() {
		for i := 0; i < 100; i++ {
			database.InsertARecord(db)
		}
	})
	for n := 0; n < b.N; n++ {
		goods := []database.Good{}
		err := sqlxDB.Select(&goods, rawSelect100SQL)
		AssertBNoError(b, err)
	}
}

func sqlxSelect(b *testing.B) {
	id := database.InsertARecord(db)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		good := database.Good{}
		err := sqlxDB.Get(&good, rawSelectSQL, id)
		AssertBNoError(b, err)
	}
}

func sqlxInsert(b *testing.B) {
	sql := rawInsertBaseSQL + rawNamedValuesSQL
	good := database.NewGood()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := sqlxDB.NamedExec(sql, good)
		AssertBNoError(b, err)
	}
}

func sqlxUpdate(b *testing.B) {
	good := database.NewGood()
	good.ID = database.InsertARecord(db)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := sqlxDB.NamedExec(rawNamedUpdateSQL, good)
		AssertBNoError(b, err)
	}
}

func sqlxDelete(b *testing.B) {
	var ids []int
	prepareWrapper(b, func() {
		ids = make([]int, b.N)
		for n := 0; n < b.N; n++ {
			ids[n] = database.InsertARecord(db)
		}
	})
	for n := 0; n < b.N; n++ {
		_, err := sqlxDB.NamedExec("DELETE FROM goods WHERE id=:id", map[string]interface{}{"id": ids[n]})
		AssertBNoError(b, err)
	}
}

func BenchmarkSQLX(b *testing.B) {
	if sqlxDB == nil {
		sqlxDB = newSQLXDB4Benchmark(b)
	}
	b.Cleanup(func() {
		if sqlxDB != nil {
			sqlxDB.Close()
		}
	})
	b.Run("select-multi", sqlxSelectMulti)
	b.Run("select", sqlxSelect)
	b.Run("insert", sqlxInsert)
	b.Run("update", sqlxUpdate)
	b.Run("delete", sqlxDelete)
}

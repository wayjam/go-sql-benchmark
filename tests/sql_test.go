package tests

import (
	"database/sql"
	"testing"

	"github.com/wayjam/go-sql-benchmark/database"
)

var (
	sqlDB *sql.DB
)

func sqlSelectMulti(b *testing.B) {
	prepareWrapper(b, func() {
		for i := 0; i < 100; i++ {
			database.InsertARecord(db)
		}
	})
	for n := 0; n < b.N; n++ {
		rows, err := sqlDB.Query(rawSelect100SQL)
		AssertBNoError(b, err)

		for rows.Next() {
			good := database.Good{}
			err := rows.Scan(&good.ID, &good.Title, &good.Description, &good.Category, &good.ThumbnailURL, &good.Price, &good.Selling, &good.Likes)
			AssertBNoError(b, err)
		}
	}
}

func sqlSelect(b *testing.B) {
	id := database.InsertARecord(db)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		good := database.Good{}
		err := sqlDB.QueryRow(rawSelectSQL, id).Scan(&good.ID, &good.Title, &good.Description, &good.Category, &good.ThumbnailURL, &good.Price, &good.Selling, &good.Likes)
		AssertBNoError(b, err)
	}
}

func sqlInsert(b *testing.B) {
	sql := rawInsertBaseSQL + rawValuesSQL
	good := database.NewGood()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := sqlDB.Exec(sql, good.Title, good.Description, good.Category, good.ThumbnailURL, good.Price, good.Selling, good.Likes)
		AssertBNoError(b, err)
	}
}

func sqlUpdate(b *testing.B) {
	toUpdateID := database.InsertARecord(db)
	good := database.NewGood()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := sqlDB.Exec(rawUpdateSQL, good.Title, good.Description, good.Category, good.ThumbnailURL, good.Price, good.Selling, good.Likes, toUpdateID)
		AssertBNoError(b, err)
	}
}

func sqlDelete(b *testing.B) {
	var ids []int
	prepareWrapper(b, func() {
		ids = make([]int, b.N)
		for n := 0; n < b.N; n++ {
			ids[n] = database.InsertARecord(db)
		}
	})
	for n := 0; n < b.N; n++ {
		_, err := sqlDB.Exec("DELETE FROM goods WHERE id=$1", ids[n])
		AssertBNoError(b, err)
	}
}

func BenchmarkSQL(b *testing.B) {
	if sqlDB == nil {
		sqlDB = newSQLDB4Benchmark(b)
	}
	b.Cleanup(func() {
		if sqlDB != nil {
			sqlDB.Close()
		}
	})
	b.Run("select-multi", sqlSelectMulti)
	b.Run("select", sqlSelect)
	b.Run("insert", sqlInsert)
	b.Run("update", sqlUpdate)
	b.Run("delete", sqlDelete)
}

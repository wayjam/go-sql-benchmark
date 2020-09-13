package tests

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"

	"github.com/wayjam/go-sql-benchmark/database"
)

var (
	pgxDB *pgx.Conn
)

func pgxSelectMulti(b *testing.B) {
	for i := 0; i < 100; i++ {
		database.InsertARecord(db)
	}
	ctx := context.Background()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		rows, err := pgxDB.Query(ctx, rawSelect100SQL)
		AssertBNoError(b, err)

		for rows.Next() {
			good := database.Good{}
			err := rows.Scan(&good.ID, &good.Title, &good.Description, &good.Category, &good.ThumbnailURL, &good.Price, &good.Selling, &good.Likes)
			AssertBNoError(b, err)
		}
	}
}

func pgxSelect(b *testing.B) {
	id := database.InsertARecord(db)
	ctx := context.Background()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		good := database.Good{}
		err := pgxDB.QueryRow(ctx, rawSelectSQL, id).Scan(&good.ID, &good.Title, &good.Description, &good.Category, &good.ThumbnailURL, &good.Price, &good.Selling, &good.Likes)
		AssertBNoError(b, err)
	}
}

func pgxInsert(b *testing.B) {
	sql := rawInsertBaseSQL + rawValuesSQL
	good := database.NewGood()
	ctx := context.Background()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := pgxDB.Exec(ctx, sql, good.Title, good.Description, good.Category, good.ThumbnailURL, good.Price, good.Selling, good.Likes)
		AssertBNoError(b, err)
	}
}

func pgxUpdate(b *testing.B) {
	toUpdateID := database.InsertARecord(db)
	good := database.NewGood()
	ctx := context.Background()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := pgxDB.Exec(ctx, rawUpdateSQL, good.Title, good.Description, good.Category, good.ThumbnailURL, good.Price, good.Selling, good.Likes, toUpdateID)
		AssertBNoError(b, err)
	}
}

func pgxDelete(b *testing.B) {
	ids := make([]int, b.N)
	for n := 0; n < b.N; n++ {
		ids[n] = database.InsertARecord(db)
	}
	ctx := context.Background()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := pgxDB.Exec(ctx, "DELETE FROM goods WHERE id=$1", ids[n])
		AssertBNoError(b, err)
	}
}

func BenchmarkPGX(b *testing.B) {
	if pgxDB == nil {
		var err error
		pgxDB, err = pgx.Connect(context.Background(), DSN)
		AssertBNoError(b, err)
	}
	b.Cleanup(func() {
		if pgxDB != nil {
			pgxDB.Close(context.Background())
		}
	})
	b.Run("select-multi", pgxSelectMulti)
	b.Run("select", pgxSelect)
	b.Run("insert", pgxInsert)
	b.Run("update", pgxUpdate)
	b.Run("delete", pgxDelete)
}

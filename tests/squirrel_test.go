package tests

import (
	"testing"

	"github.com/wayjam/go-sql-benchmark/database"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var (
	sq       squirrel.StatementBuilderType
	sqSQLXDB *sqlx.DB
)

func squirrelSelectMulti(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		query := sq.Select("*").From("goods").Limit(100).RunWith(sqSQLXDB)

		rows, err := query.Query()
		AssertBNoError(b, err)

		for rows.Next() {
			good := database.Good{}
			err := rows.Scan(&good.ID, &good.Title, &good.Description, &good.Category, &good.ThumbnailURL, &good.Price, &good.Selling, &good.Likes)
			AssertBNoError(b, err)
		}
	}
}

func squirrelSelectMultiWithSQLX(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sql, args, err := sq.Select("*").From("goods").Limit(100).ToSql()
		AssertBNoError(b, err)

		goods := []database.Good{}

		err = sqSQLXDB.Select(&goods, sql, args...)
		AssertBNoError(b, err)
	}
}

func squirrelSelect(b *testing.B) {
	id := database.InsertARecord(db)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		good := database.Good{}
		err := sq.
			Select("*").
			From("goods").
			Where(squirrel.Eq{"id": id}).
			RunWith(sqSQLXDB).
			QueryRow().
			Scan(&good.ID, &good.Title, &good.Description, &good.Category, &good.ThumbnailURL, &good.Price, &good.Selling, &good.Likes)
		AssertBNoError(b, err)
	}
}

func squirrelSelectWithSQLX(b *testing.B) {
	id := database.InsertARecord(db)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		good := database.Good{}
		sql, args, err := sq.
			Select("*").
			From("goods").
			Where(squirrel.Eq{"id": id}).ToSql()
		AssertBNoError(b, err)
		err = sqSQLXDB.Get(&good, sql, args...)
		AssertBNoError(b, err)
	}
}

func squirrelBasicInsertBuilder(good *database.Good) squirrel.InsertBuilder {
	return sq.
		Insert("goods").
		Columns("title", "description", "category", "thumbnail", "price", "selling", "likes").
		Values(good.Title, good.Description, good.Category, good.ThumbnailURL, good.Price, good.Selling, good.Likes)
}

func squirrelInsert(b *testing.B) {
	good := database.NewGood()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := squirrelBasicInsertBuilder(good).RunWith(sqSQLXDB).Exec()
		AssertBNoError(b, err)
	}
}

func squirrelInsertWithSQLX(b *testing.B) {
	good := database.NewGood()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sql, args, err := squirrelBasicInsertBuilder(good).ToSql()
		AssertBNoError(b, err)

		_, err = sqSQLXDB.Exec(sql, args...)
		AssertBNoError(b, err)
	}
}

func squirrelBasicUpdateBuilder(good *database.Good) squirrel.UpdateBuilder {
	return sq.Update("goods").
		Set("title", good.Title).
		Set("description", good.Description).
		Set("category", good.Category).
		Set("thumbnail", good.ThumbnailURL).
		Set("price", good.Price).
		Set("selling", good.Selling).
		Set("likes", good.Likes).
		Where(squirrel.Eq{"id": good.ID})
}

func squirrelUpdate(b *testing.B) {
	good := database.NewGood()
	good.ID = database.InsertARecord(db)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := squirrelBasicUpdateBuilder(good).RunWith(sqSQLXDB).Exec()
		AssertBNoError(b, err)
	}
}

func squirrelUpdateWithSQLX(b *testing.B) {
	good := database.NewGood()
	good.ID = database.InsertARecord(db)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sql, args, err := squirrelBasicUpdateBuilder(good).ToSql()

		_, err = sqSQLXDB.Exec(sql, args...)
		AssertBNoError(b, err)
	}
}

func squirrelDeleteWithSQLX(b *testing.B) {
	var ids []int
	prepareWrapper(b, func() {
		ids = make([]int, b.N)
		for n := 0; n < b.N; n++ {
			ids[n] = database.InsertARecord(db)
		}
	})
	for n := 0; n < b.N; n++ {
		sql, args, err := sq.Delete("goods").Where(squirrel.Eq{"id": ids[n]}).ToSql()
		AssertBNoError(b, err)
		sqSQLXDB.Exec(sql, args...)
		AssertBNoError(b, err)
	}
}

func BenchmarkSquirrel(b *testing.B) {
	sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	if sqSQLXDB == nil {
		sqSQLXDB = newSQLXDB4Benchmark(b)
	}

	b.Cleanup(func() {
		if sqSQLXDB != nil {
			sqSQLXDB.Close()
		}
	})

	b.Run("select-multi", squirrelSelectMulti)
	b.Run("select-multi-with-sqlx", squirrelSelectMultiWithSQLX)
	b.Run("select", squirrelSelect)
	b.Run("select-with-sqlx", squirrelSelectWithSQLX)
	b.Run("insert", squirrelInsert)
	b.Run("insert-with-sqlx", squirrelInsertWithSQLX)
	b.Run("update", squirrelUpdate)
	b.Run("update-with-sqlx", squirrelUpdateWithSQLX)
	b.Run("delete-with-sqlx", squirrelDeleteWithSQLX)
}

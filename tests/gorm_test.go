package tests

import (
	"testing"

	"github.com/wayjam/go-sql-benchmark/database"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	gormDB *gorm.DB
)

func gormSelectMulti(b *testing.B) {
	for i := 0; i < 100; i++ {
		database.InsertARecord(db)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var goods []database.Good
		result := gormDB.Limit(100).Find(&goods)
		AssertBNoError(b, result.Error)
	}
}

func gormSelect(b *testing.B) {
	id := database.InsertARecord(db)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		good := new(database.Good)
		result := gormDB.Find(good, id)
		AssertBNoError(b, result.Error)
	}
}

func gormInsert(b *testing.B) {
	good := database.NewGood()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result := gormDB.Omit("id").Create(good)
		AssertBNoError(b, result.Error)
	}
}

func gormUpdate(b *testing.B) {
	toUpdateID := database.InsertARecord(db)
	toUpdateModel := database.Good{ID: toUpdateID}
	good := database.NewGood()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result := gormDB.Model(&toUpdateModel).Updates(good)
		AssertBNoError(b, result.Error)
	}
}

func gormDelete(b *testing.B) {
	var ids []int
	prepareWrapper(b, func() {
		ids = make([]int, b.N)
		for n := 0; n < b.N; n++ {
			ids[n] = database.InsertARecord(db)
		}
	})
	for n := 0; n < b.N; n++ {
		result := gormDB.Delete(&database.Good{}, ids[n])
		AssertBNoError(b, result.Error)
	}
}

func BenchmarkGORM(b *testing.B) {
	if gormDB == nil {
		var err error
		// https://github.com/go-gorm/postgres
		gormDB, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  DSN,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
		AssertBNoError(b, err)

		stddb, err := gormDB.DB()
		AssertBNoError(b, err)

		err = stddb.Ping()
		AssertBNoError(b, err)
	}
	b.Cleanup(func() {
		if gormDB != nil {
			stddb, _ := gormDB.DB()
			if stddb != nil {
				stddb.Close()
			}
		}
	})
	b.Run("select-multi", gormSelectMulti)
	b.Run("select", gormSelect)
	b.Run("insert", gormInsert)
	b.Run("update", gormUpdate)
	b.Run("delete", gormDelete)
}

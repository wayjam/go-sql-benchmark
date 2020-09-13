package database

import (
	"database/sql"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Good struct {
	ID           int    `db:"id" gorm:"primaryKey"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	Category     string `db:"category"`
	ThumbnailURL string `db:"thumbnail" gorm:"column:thumbnail"`
	Price        int    `db:"price"`
	Selling      bool   `db:"selling"`
	Likes        int64  `db:"likes"`
}

var (
	counter int = 0
)

func NewGood() *Good {
	strCount := strconv.Itoa(counter)
	counter += 1
	return &Good{
		Title:        "good-" + strCount,
		Description:  "a-desc-string" + strCount,
		Category:     "cate-" + strconv.Itoa(rand.Intn(20)),
		ThumbnailURL: "https://example.com/good/" + strCount + "/thumbnail",
		Price:        rand.Intn(9999),
		Selling:      rand.Float32() < 0.5,
		Likes:        rand.Int63n(int64(999)),
	}
}

func InsertARecord(db *sql.DB) (id int) {
	p := NewGood()
	row := db.QueryRow(
		`INSERT INTO goods (title, description, category, thumbnail, price, selling, likes) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		p.Title,
		p.Description,
		p.Category,
		p.ThumbnailURL,
		p.Price,
		p.Selling,
		p.Likes,
	)

	row.Scan(&id)
	return
}

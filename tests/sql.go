package tests

var (
	rawInsertBaseSQL  = `INSERT INTO goods (title, description, category, thumbnail, price, selling, likes) VALUES `
	rawValuesSQL      = `($1, $2, $3, $4, $5, $6, $7)`
	rawNamedValuesSQL = `(:title, :description, :category, :thumbnail, :price, :selling, :likes)`
	rawUpdateSQL      = `UPDATE goods SET title=$1, description=$2, category=$3, thumbnail=$4, price=$5, selling=$6, likes=$7 WHERE id=$8`
	rawNamedUpdateSQL = `UPDATE goods SET title=:title, description=:description, category=:category, thumbnail=:thumbnail, price=:price, selling=:selling, likes=:likes WHERE id=:id`
	rawSelect100SQL   = `SELECT * FROM goods WHERE id>0 LIMIT 100`
	rawSelectSQL      = `SELECT * FROM goods WHERE id=$1`
)

package database

import (
	"database/sql"
)

// init sqls
var (
	InitSQLS []string = []string{
		`DROP TABLE IF EXISTS goods;`,
		`CREATE TABLE goods(
			id SERIAL NOT NULL,
			title varchar(255) NOT NULL,
			description text,
			category varchar(255),
			thumbnail varchar(255),
			price integer NOT NULL,
			selling boolean default false,
			likes bigint default 0,
			CONSTRAINT goods_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,
	}
)

func InitDB(db *sql.DB) (err error) {
	err = db.Ping()
	if err != nil {
		return
	}

	for _, sql := range InitSQLS {
		_, err = db.Exec(sql)
		if err != nil {
			return
		}
	}
	return nil
}

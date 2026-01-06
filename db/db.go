package db

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func ConnectDB() (*sql.DB, error) {
	connString :=
		"server=localhost;" +
			"user id=sa;" +
			"password=123456;" +
			"database=ShortLink"
	return sql.Open("sqlserver", connString)
}

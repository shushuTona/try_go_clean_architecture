package drivers

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func CreateDB() (*sql.DB, error) {
	var jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	var c = mysql.Config{
		DBName:    "db",
		User:      "root",
		Passwd:    "root",
		Addr:      "db:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

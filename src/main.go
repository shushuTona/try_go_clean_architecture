package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"

	"try_go_clean_architecture/entities"
	"try_go_clean_architecture/repository/mysql/githubRepo"
)

func main() {
	var jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic("db error.")
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
		panic("db open error.")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	githubRepoMysqlRepository := githubRepo.NewGithubRepoMysqlRepository(db)

	var grl = []entities.GithubRepo{
		{
			ID:        400,
			Name:      "test repo 3",
			UpdatedAt: "2023-03-05 09:00:00",
		},
	}
	err = githubRepoMysqlRepository.InsertRepos(grl)

	fmt.Println(err.Error())
}

package githubRepo

import (
	"database/sql"
	"try_go_clean_architecture/entities"
)

type IMysqlGateway interface {
	GetRepos() ([]entities.GithubRepo, error)
	InsertRepos(grlist []entities.GithubRepo) error
}

type MysqlGatewayFactory = func(db *sql.DB) IMysqlGateway

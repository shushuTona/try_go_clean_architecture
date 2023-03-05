package mysql

import "try_go_clean_architecture/entities"

type IMysqlRepository interface {
	GetRepos() []entities.GithubRepo
}

package githubRepo

import "try_go_clean_architecture/entities"

type IMysqlRepository interface {
	GetRepos() ([]entities.GithubRepo, error)
	InsertRepos(grlist []entities.GithubRepo) error
}

package githubRepo

import (
	"database/sql"
	"try_go_clean_architecture/entities"
)

type mockGithubRepoMysqlGateway struct {
	db *sql.DB
}

func NewMockGithubRepoMysqlGateway() IMysqlGateway {
	var db *sql.DB
	return &mockGithubRepoMysqlGateway{
		db: db,
	}
}

func (gr *mockGithubRepoMysqlGateway) GetRepos() ([]entities.GithubRepo, error) {
	var githubRepoList = []entities.GithubRepo{
		{
			ID:        100,
			Name:      "mock test repo name 1",
			UpdatedAt: "2023-01-01 00:00:00",
		},
		{
			ID:        200,
			Name:      "mock test repo name 2",
			UpdatedAt: "2023-02-01 00:00:00",
		},
		{
			ID:        300,
			Name:      "mock test repo name 3",
			UpdatedAt: "2023-03-01 00:00:00",
		},
	}

	return githubRepoList, nil
}

func (gr *mockGithubRepoMysqlGateway) InsertRepos(grlist []entities.GithubRepo) error {
	return nil
}

package mysql

import (
	"database/sql"
	"try_go_clean_architecture/entities"
)

type githubRepoMysqlRepository struct {
	db *sql.DB
}

func NewGithubRepoMysqlRepository(db *sql.DB) IMysqlRepository {
	return &githubRepoMysqlRepository{
		db: db,
	}
}

func (gr *githubRepoMysqlRepository) GetRepos() []entities.GithubRepo {
	query := `SELECT id, name, updated_at FROM repos`
	stmt, err := gr.db.Prepare(query)
	if err != nil {
		panic("db prepare error")
	}

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var githubRepoList []entities.GithubRepo
	for rows.Next() {
		githubRepo := entities.GithubRepo{}

		err = rows.Scan(
			&githubRepo.ID,
			&githubRepo.Name,
			&githubRepo.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}

		githubRepoList = append(githubRepoList, githubRepo)
	}

	return githubRepoList
}

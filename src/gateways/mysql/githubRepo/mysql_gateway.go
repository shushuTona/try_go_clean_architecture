package githubRepo

import (
	"database/sql"
	"fmt"
	"try_go_clean_architecture/entities"
)

type githubRepoMysqlGateway struct {
	db *sql.DB
}

func NewGithubRepoMysqlGateway(db *sql.DB) IMysqlGateway {
	return &githubRepoMysqlGateway{
		db: db,
	}
}

func (gr *githubRepoMysqlGateway) GetRepos() ([]entities.GithubRepo, error) {
	query := `SELECT id, name, updated_at FROM repos`
	stmt, err := gr.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
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
			return nil, err
		}

		githubRepoList = append(githubRepoList, githubRepo)
	}

	return githubRepoList, nil
}

func (gr *githubRepoMysqlGateway) InsertRepos(grlist []entities.GithubRepo) error {
	// starts a transaction
	tx, err := gr.db.Begin()
	if err != nil {
		fmt.Println("gr.db.Begin()")
		return err
	}

	// create preoare
	stmt, err := tx.Prepare("INSERT INTO repos (id, name, updated_at) VALUES( ?, ?, ? )")
	if err != nil {
		tx.Rollback()
		fmt.Println("create stmt")
		return err
	}
	defer stmt.Close()

	// insert repository list
	for _, githubRepo := range grlist {
		if _, err := stmt.Exec(githubRepo.ID, githubRepo.Name, githubRepo.UpdatedAt); err != nil {
			tx.Rollback()
			fmt.Println("range loop")
			return err
		}
	}

	tx.Commit()

	return nil
}

package githubRepo

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"try_go_clean_architecture/entities"
)

func TestGetRepos(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// prepare mock
	rows := sqlmock.NewRows([]string{"id", "name", "updated_at"}).AddRow(100, "test repo 100", "2023-01-01 00:00:00")
	query := "SELECT id, name, updated_at FROM repos"
	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WillReturnRows(rows)

	githubRepoMysqlRepository := NewGithubRepoMysqlRepository(db)
	githubRepoList, err := githubRepoMysqlRepository.GetRepos()

	if err != nil {
		t.Error("err is not nil")
	}
	if len(githubRepoList) != 1 {
		t.Error("githubRepoList length is not 1")
	}
	if githubRepoList[0].ID != 100 {
		t.Error("repo ID is not 100")
	}
	if githubRepoList[0].Name != "test repo 100" {
		t.Error("repo Name is not match")
	}
	if githubRepoList[0].UpdatedAt != "2023-01-01 00:00:00" {
		t.Error("repo UpdatedAt is not match")
	}
}

func TestInsertRepos(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO repos (id, name, updated_at) VALUES( ?, ?, ? )")).ExpectQuery()
	//  .WithArgs(999, "test name", "2023-01-01 00:00:00").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	githubRepoMysqlRepository := NewGithubRepoMysqlRepository(db)
	var grlist = []entities.GithubRepo{
		{
			ID:        999,
			Name:      "test name",
			UpdatedAt: "2023-01-01 00:00:00",
		},
	}
	err = githubRepoMysqlRepository.InsertRepos(grlist)
	if err != nil {
		t.Error(err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err.Error())
	}
}

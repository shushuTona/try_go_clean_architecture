package usecases

import (
	"testing"
	githubRepoGateway "try_go_clean_architecture/gateways/mysql/githubRepo"
	"try_go_clean_architecture/presenters"
)

func TestShowRepos(t *testing.T) {
	var mockGateway = githubRepoGateway.NewMockGithubRepoMysqlGateway()
	githubRepoList, err := mockGateway.GetRepos()
	if err != nil {
		t.Errorf("mockGateway.GetRepos err is not nil : %s", err.Error())
	}

	var mockPresenter = presenters.NewMockGithubReposPresenter()
	err = mockPresenter.Display(githubRepoList)
	if err != nil {
		t.Errorf("mockPresenter.Display err is not nil : %s", err.Error())
	}
}

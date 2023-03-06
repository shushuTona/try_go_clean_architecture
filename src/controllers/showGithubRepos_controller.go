package controllers

import (
	"database/sql"
	"io"
	githubRepoGateway "try_go_clean_architecture/gateways/mysql/githubRepo"
	"try_go_clean_architecture/presenters"
	"try_go_clean_architecture/usecases"
)

type GithubReposController struct {
	GatewayFactory   githubRepoGateway.MysqlGatewayFactory
	PresenterFactory presenters.GithubReposPresenterFactory
	UsecaseFactory   usecases.DisplayGithubReposUseCaseFactory
}

func (grc *GithubReposController) ShowRepos(db *sql.DB, dist io.Writer) error {
	var mysqlGateway = grc.GatewayFactory(db)
	var githubReposPresenter = grc.PresenterFactory(dist)
	var displayGithubReposUseCase = grc.UsecaseFactory(mysqlGateway, githubReposPresenter)

	var err = displayGithubReposUseCase.ShowRepos()
	if err != nil {
		return err
	}

	return nil
}

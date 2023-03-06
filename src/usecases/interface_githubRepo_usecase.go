package usecases

import (
	githubRepoGateway "try_go_clean_architecture/gateways/mysql/githubRepo"
	"try_go_clean_architecture/presenters"
)

type IDisplayGithubReposUseCase interface {
	ShowRepos() error
}

type DisplayGithubReposUseCaseFactory = func(grg githubRepoGateway.IMysqlGateway, grp presenters.IGithubReposPresenter) IDisplayGithubReposUseCase

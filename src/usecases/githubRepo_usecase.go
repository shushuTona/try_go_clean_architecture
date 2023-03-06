package usecases

import (
	githubRepoGateway "try_go_clean_architecture/gateways/mysql/githubRepo"
	"try_go_clean_architecture/presenters"
)

type displayGithubReposUseCase struct {
	grg githubRepoGateway.IMysqlGateway
	grp presenters.IGithubReposPresenter
}

func NewDisplayGithubReposUseCase(grg githubRepoGateway.IMysqlGateway, grp presenters.IGithubReposPresenter) IDisplayGithubReposUseCase {
	return &displayGithubReposUseCase{
		grg: grg,
		grp: grp,
	}
}

func (dgr *displayGithubReposUseCase) ShowRepos() error {
	githubRepoList, err := dgr.grg.GetRepos()
	if err != nil {
		return err
	}

	err = dgr.grp.Display(githubRepoList)
	if err != nil {
		return err
	}

	return nil
}

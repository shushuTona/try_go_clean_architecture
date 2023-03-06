package presenters

import (
	"io"
	"try_go_clean_architecture/entities"
)

type IGithubReposPresenter interface {
	Display(grl []entities.GithubRepo) error
}

type GithubReposPresenterFactory = func(dist io.Writer) IGithubReposPresenter

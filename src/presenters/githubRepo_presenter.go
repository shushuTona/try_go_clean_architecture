package presenters

import (
	"io"
	"try_go_clean_architecture/entities"
)

type githubReposPresenter struct {
	dist io.Writer
}

func NewGithubReposPresenter(dist io.Writer) IGithubReposPresenter {
	return &githubReposPresenter{
		dist: dist,
	}
}

func (grp *githubReposPresenter) Display(grl []entities.GithubRepo) error {
	var githubReposNameStr string
	for _, gr := range grl {
		githubReposNameStr = githubReposNameStr + gr.Name + "\n"
	}

	var _, err = grp.dist.Write([]byte(githubReposNameStr))
	if err != nil {
		return err
	}

	return nil
}

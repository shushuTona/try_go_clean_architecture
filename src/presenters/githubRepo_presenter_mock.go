package presenters

import (
	"io"
	"os"
	"try_go_clean_architecture/entities"
)

type mockGithubReposPresenter struct {
	dist io.Writer
}

func NewMockGithubReposPresenter() IGithubReposPresenter {
	return &mockGithubReposPresenter{
		dist: os.Stdout,
	}
}

func (grp *mockGithubReposPresenter) Display(grl []entities.GithubRepo) error {
	var githubReposNameStr string
	for _, gr := range grl {
		githubReposNameStr = githubReposNameStr + gr.Name + "\n"
	}

	grp.dist.Write([]byte(githubReposNameStr))

	return nil
}

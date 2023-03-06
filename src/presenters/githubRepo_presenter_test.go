package presenters

import (
	"bytes"
	"testing"
	"try_go_clean_architecture/entities"
)

func TestDisplay(t *testing.T) {
	var dist = bytes.Buffer{}
	var githubReposPresenter = NewGithubReposPresenter(&dist)

	var grl = []entities.GithubRepo{
		{
			ID:        100,
			Name:      "test repo name1",
			UpdatedAt: "2023-01-01 00:00:00",
		},
		{
			ID:        200,
			Name:      "test repo name2",
			UpdatedAt: "2023-01-01 00:00:00",
		},
		{
			ID:        300,
			Name:      "test repo name3",
			UpdatedAt: "2023-01-01 00:00:00",
		},
	}
	var err = githubReposPresenter.Display(grl)
	if err != nil {
		t.Error(err.Error())
	}

	var expect = "test repo name1\ntest repo name2\ntest repo name3\n"
	if dist.String() != expect {
		t.Errorf("Error get: %s, want: %s", dist.String(), expect)
	}
}

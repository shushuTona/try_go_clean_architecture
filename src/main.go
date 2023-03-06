package main

import (
	"fmt"
	"os"

	"try_go_clean_architecture/controllers"
	"try_go_clean_architecture/drivers"
	githubRepoGateway "try_go_clean_architecture/gateways/mysql/githubRepo"
	"try_go_clean_architecture/presenters"
	"try_go_clean_architecture/usecases"
)

func main() {
	var db, err = drivers.CreateDB()
	if err != nil {
		panic(err.Error())
	}
	// var dist = os.Stdout
	dist, err := os.OpenFile("./text.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err.Error())
	}

	var githubReposController = controllers.GithubReposController{
		GatewayFactory:   githubRepoGateway.NewGithubRepoMysqlGateway,
		PresenterFactory: presenters.NewGithubReposPresenter,
		UsecaseFactory:   usecases.NewDisplayGithubReposUseCase,
	}

	err = githubReposController.ShowRepos(db, dist)

	if err != nil {
		fmt.Println(err.Error())
	}
}

package app

import (
	"golang_mvc/git_repo/src/api/controllers/polo"
	"golang_mvc/git_repo/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)

}

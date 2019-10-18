package main

import (
	"os"

	"github.com/multi-repo-docs-versioning/docs-manager/pkg/repository"
	utils "github.com/multi-repo-docs-versioning/docs-manager/pkg/utils"
	"gopkg.in/src-d/go-git.v4"
)

type TagName int

const (
	Master TagName = iota
	Experimental
)

func (t TagName) String() string {
	return [...]string{"master", "experimental"}[t]
}

func versionHandler(config *utils.DocsConfig) {
	versions := config.GetDocsYamlConfig().Versions
	for _, val := range versions {
		switch val.Ver {
		case Master.String():
			repos := val.Repos
			for _, repo := range repos {
				path := os.Args[2] + repo.Name
				err := os.MkdirAll(path, 0700)

				if err != nil {
					panic(err)
				}
				cloneOptions := git.CloneOptions{
					URL:      repo.URL,
					Tags:     git.AllTags,
					Progress: os.Stdout,
				}
				repo := repository.New().
					SetCloneOptions(cloneOptions).
					SetTagName(repo.TagName).
					SetPath(path).
					Build()

				err = repo.Clone()
				utils.CheckIfError(err)
			}
		case Experimental.String():
			repos := val.Repos
			for _, repo := range repos {
				path := os.Args[2] + "experimental/" + repo.Name
				if _, err := os.Stat(path); os.IsNotExist(err) {
					err = os.MkdirAll(path, 0700)
					utils.CheckIfError(err)
				}

				cloneOptions := git.CloneOptions{
					URL:      repo.URL,
					Tags:     git.AllTags,
					Progress: os.Stdout,
				}
				repo := repository.New().
					SetCloneOptions(cloneOptions).
					SetTagName(repo.TagName).
					SetPath(path).
					Build()

				err := repo.Clone()
				utils.CheckIfError(err)
			}
		}
	}
}

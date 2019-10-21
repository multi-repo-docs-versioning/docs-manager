package main

import (
	"os"

	"github.com/multi-repo-docs-versioning/docs-manager/pkg/repository"
	utils "github.com/multi-repo-docs-versioning/docs-manager/pkg/utils"
	"gopkg.in/src-d/go-git.v4"
)

type TagName int

const (
	Latest TagName = iota
	Experimental
)

func (t TagName) String() string {
	return [...]string{"v1.0", "v0.0"}[t]
}

func versionHandler(config *utils.DocsConfig) {
	versions := config.GetDocsYamlConfig().Versions
	err := os.MkdirAll("./site", 0755)
	utils.CheckIfError(err)
	versionsArray := make([]string, len(versions))

	for index, val := range versions {
		versionsArray[index] = val.Ver
	}

	for _, val := range versions {
		switch val.Ver {
		case Latest.String():
			repos := val.Repos
			for _, repo := range repos {
				path := os.Args[2] + repo.Name
				err := os.MkdirAll(path, 0755)
				utils.CheckIfError(err)
				cloneOptions := git.CloneOptions{
					URL:      repo.URL,
					Tags:     git.AllTags,
					Progress: os.Stdout,
				}
				gitRepo := repository.New().
					SetCloneOptions(cloneOptions).
					SetTagName(repo.TagName).
					SetPath(path).
					Build()

				err = gitRepo.Clone()
				utils.CheckIfError(err)
				err = utils.RemoveContents(path)
				utils.CheckIfError(err)
			}
			build(versionsArray, Latest.String())
		/*case Experimental.String():
		repos := val.Repos
		docsDir := os.Args[2] + Experimental.String() + "/"
		for _, repo := range repos {
			path := docsDir + repo.Name
			if _, err := os.Stat(path); os.IsNotExist(err) {
				err = os.MkdirAll(path, 0755)
				utils.CheckIfError(err)
			}

			cloneOptions := git.CloneOptions{
				URL:      repo.URL,
				Tags:     git.AllTags,
				Progress: os.Stdout,
			}
			gitRepo := repository.New().
				SetCloneOptions(cloneOptions).
				SetTagName(repo.TagName).
				SetPath(path).
				Build()

			err := gitRepo.Clone()
			utils.CheckIfError(err)
			err = removeContents(path)
			utils.CheckIfError(err)
		}
		build(versionsArray, Experimental.String())*/
		default:
			repos := val.Repos
			docsDir := os.Args[2] + val.Ver + "/"
			for _, repo := range repos {
				path := docsDir + repo.Name
				if _, err := os.Stat(path); os.IsNotExist(err) {
					err = os.MkdirAll(path, 0755)
					utils.CheckIfError(err)
				}

				cloneOptions := git.CloneOptions{
					URL:      repo.URL,
					Tags:     git.AllTags,
					Progress: os.Stdout,
				}
				gitRepo := repository.New().
					SetCloneOptions(cloneOptions).
					SetTagName(repo.TagName).
					SetPath(path).
					Build()

				err := gitRepo.Clone()
				utils.CheckIfError(err)
				err = gitRepo.CheckOutTag()
				utils.CheckIfError(err)
				err = utils.RemoveContents(path)
				utils.CheckIfError(err)
			}
			build(versionsArray, val.Ver)
		}
	}
}

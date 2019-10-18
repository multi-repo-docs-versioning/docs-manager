package main

import (
	"os"

	"github.com/multi-repo-docs-versioning/docs-manager/pkg/repository"
	"gopkg.in/src-d/go-git.v4"

	utils "github.com/multi-repo-docs-versioning/docs-manager/pkg/utils"
)

func main() {
	config := utils.NewDocsConfig(os.Args[1])
	err := config.Parse()
	utils.CheckIfError(err)
	versions := config.GetDocsYamlConfig().Versions
	for _, val := range versions {
		switch val.Ver {
		case "master":
			repos := val.Repos
			for _, repo := range repos {
				path := os.Args[2] + repo.Name
				err = os.Mkdir(path, 0700)

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

		}
	}

}

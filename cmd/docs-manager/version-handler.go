package main

import (
	"os"
	"path/filepath"

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

func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		if name != "docs" {
			err = os.RemoveAll(filepath.Join(dir, name))
			if err != nil {
				return err
			}
		}
	}
	return nil
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
				err = removeContents(path)
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
				err = removeContents(path)
				utils.CheckIfError(err)
			}
		default:
			repos := val.Repos
			for _, repo := range repos {
				path := os.Args[2] + val.Ver + "/" + repo.Name
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
				err = repo.CheckOutTag()
				utils.CheckIfError(err)
				err = removeContents(path)
				utils.CheckIfError(err)
			}

		}
	}
}

package build

import (
	"os"

	"github.com/multi-repo-docs-versioning/docs-manager/pkg/common"

	"github.com/multi-repo-docs-versioning/docs-manager/pkg/repository"
	utils "github.com/multi-repo-docs-versioning/docs-manager/pkg/utils"
	"gopkg.in/src-d/go-git.v4"
)

type TagName int

const (
	Latest TagName = iota
)

func (t TagName) String() string {
	return [...]string{"v1.0"}[t]
}

// VersionHandler handle different versions of docs
func VersionHandler(config *utils.DocsConfig) {
	versions := config.GetDocsYamlConfig().Versions
	err := os.MkdirAll(common.SiteDirName, common.PermissionMode)
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
				err := os.MkdirAll(path, common.PermissionMode)
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
		default:
			repos := val.Repos
			docsDir := os.Args[2] + val.Ver + "/"
			for _, repo := range repos {
				path := docsDir + repo.Name
				if _, err := os.Stat(path); os.IsNotExist(err) {
					err = os.MkdirAll(path, common.PermissionMode)
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

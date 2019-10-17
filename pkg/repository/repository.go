package repository

import (
	utils "github.com/multi-repo-docs-versioning/docs-manager/pkg/utils"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type Repository struct {
	path         string
	cloneOptions git.CloneOptions
	gitRepo      *git.Repository
	tagName      string
}

type RepositoryBuilder interface {
	SetPath(string) RepositoryBuilder
	SetCloneOptions(git.CloneOptions) RepositoryBuilder
	SetGitRepo(*git.Repository) RepositoryBuilder
	SetTagName(string) RepositoryBuilder
	Build() Repository
}

func New() RepositoryBuilder {
	return &Repository{
		path:         "/tmp/test",
		cloneOptions: git.CloneOptions{},
		gitRepo:      &git.Repository{},
		tagName:      "master",
	}
}

func (repo *Repository) Build() Repository {
	return Repository{
		path:         repo.path,
		cloneOptions: repo.cloneOptions,
		gitRepo:      repo.gitRepo,
		tagName:      repo.tagName,
	}

}

// SetPath set repo path
func (repo *Repository) SetPath(path string) RepositoryBuilder {
	repo.path = path
	return repo
}

// SetCloneOptions set git clone options
func (repo *Repository) SetCloneOptions(cloneOptions git.CloneOptions) RepositoryBuilder {
	repo.cloneOptions = cloneOptions
	return repo
}

// SetGitRepo set git repo
func (repo *Repository) SetGitRepo(gitRepo *git.Repository) RepositoryBuilder {
	repo.gitRepo = gitRepo
	return repo
}

// SetTagName set the git repo tag name
func (repo *Repository) SetTagName(tagName string) RepositoryBuilder {
	repo.tagName = tagName
	return repo
}

//Clone clones a repo based on a given url and a path
func (repo *Repository) Clone() error {
	utils.Info("git clone %s", repo.cloneOptions.URL)
	r, err := git.PlainClone(repo.path, false, &repo.cloneOptions)
	repo.SetGitRepo(r)
	return err
}

// GetTag get a repo tag reference based on a given tag name
func (repo *Repository) GetTag() (*plumbing.Reference, error) {
	_, err := repo.gitRepo.Worktree()
	if err != nil {
		return nil, err
	}

	tagRepo, err := repo.gitRepo.Tag(repo.tagName)
	if err != nil {
		return nil, err
	}

	return tagRepo, err
}

// CheckOut check out a
func (repo *Repository) CheckOutTag() error {
	w, err := repo.gitRepo.Worktree()
	if err != nil {
		return err
	}

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName("refs/tags/" + repo.tagName),
	})

	return err

}

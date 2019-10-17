package main

import (
	"os"

	utils "github.com/multi-repo-docs-versioning/docs-manager/pkg/utils"
	"gopkg.in/src-d/go-git.v4"
)

func main() {
	// Clone the given repository to the given directory
	utils.Info("git clone https://github.com/src-d/go-git")

	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/src-d/go-git",
		Progress: os.Stdout,
	})

	utils.CheckIfError(err)

}

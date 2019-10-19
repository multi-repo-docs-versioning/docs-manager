package main

import (
	"os"

	utils "github.com/multi-repo-docs-versioning/docs-manager/pkg/utils"
)

func main() {
	config := utils.NewDocsConfig(os.Args[1])
	err := config.Parse()
	utils.CheckIfError(err)
	versionHandler(config)

}

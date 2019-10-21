package main

import (
	"os"

	"github.com/multi-repo-docs-versioning/docs-manager/pkg/menu"
	"github.com/multi-repo-docs-versioning/docs-manager/pkg/types"

	"github.com/multi-repo-docs-versioning/docs-manager/pkg/manifest"
	utils "github.com/multi-repo-docs-versioning/docs-manager/pkg/utils"
)

const (
	MkdocsConfig   = "mkdocs.yml"
	SiteDirName    = "./site/"
	PermissionMode = 0755
)

// build build docs website according to the given list of versions
func build(versions []string, tagName string) {
	manif, _ := manifest.Read(os.Args[3])
	var docsDir string

	manifestPath := MkdocsConfig
	var siteDir string
	if tagName == Latest.String() {
		docsDir = os.Args[2]
		siteDir = SiteDirName
	} else {
		docsDir = os.Args[2] + tagName + "/"
		siteDir = SiteDirName + tagName + ""
		utils.RunCommand("cp", "./content/README.md", docsDir)
		utils.RunCommand("cp", "-r", "./content/images", docsDir)
		utils.RunCommand("cp", "-r", "./content/developers", docsDir)
	}

	manif["docs_dir"] = docsDir
	err := manifest.Write(manifestPath, manif)
	utils.CheckIfError(err)
	menuConfig := types.MenuFiles{
		JsFile: os.Args[4],
	}
	menuContent := menu.GetTemplateContent(&menuConfig)

	versionsInfo := types.VersionsInformation{
		Current:      tagName,
		Latest:       Latest.String(),
		Experimental: Experimental.String(),
		CurrentPath:  docsDir,
	}

	err = menu.Build(versionsInfo, versions, menuContent)
	utils.CheckIfError(err)

	utils.RunCommand("mkdocs", "build", "--site-dir", siteDir, "-q")
	manif["docs_dir"] = "content"
	err = manifest.Write(manifestPath, manif)
	utils.CheckIfError(err)

}

package main

import (
	"os"
	"os/exec"

	"github.com/multi-repo-docs-versioning/docs-manager/pkg/menu"
	"github.com/multi-repo-docs-versioning/docs-manager/pkg/types"

	"github.com/multi-repo-docs-versioning/docs-manager/pkg/manifest"
	utils "github.com/multi-repo-docs-versioning/docs-manager/pkg/utils"
)

const (
	MkdocsConfig = "mkdocs.yml"
	SiteDirName  = "./site/"
)

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
		cmd := exec.Command("cp", "./content/README.md", docsDir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		utils.CheckIfError(err)
		cmd = exec.Command("cp", "-r", "./content/images", docsDir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		utils.CheckIfError(err)
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

	cmd := exec.Command("mkdocs", "build", "--site-dir", siteDir, "-q")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	utils.CheckIfError(err)
	manif["docs_dir"] = "content"
	err = manifest.Write(manifestPath, manif)
	utils.CheckIfError(err)

}

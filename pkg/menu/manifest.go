package menu

import "github.com/multi-repo-docs-versioning/docs-manager/pkg/manifest"

func editManifest(manif map[string]interface{}, versionJsFile string, versionCSSFile string) {
	// Append menu JS file
	manifest.AppendExtraJs(manif, versionJsFile)

	// Append menu CSS file
	manifest.AppendExtraCSS(manif, versionCSSFile)

	// reset site URL
	manif["site_url"] = ""
}

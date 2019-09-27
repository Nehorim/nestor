package typescript

import (
	"encoding/json"
)
import (
	"github.com/Nehorim/nestor/commons"
)

type PackageRepository struct {
	RepoType string `json:"type"`
	Url      string `json:"url"`
}
type PackageBugs struct {
	Url string `json:"url"`
}
type PackageScripts struct {
	Test      string `json:"test"`
	TestWatch string `json:"test:watch"`
	Start     string `json:"start"`
	Build     string `json:"build"`
	Buildlive string `json:"build:live"`
	Lint      string `json:"lint"`
	LintFix   string `json:"lint:fix"`
}
type Package struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Description string            `json:"description"`
	Main        string            `json:"main"`
	Author      string            `json:"author"`
	License     string            `json:"license"`
	Homepage    string            `json:"homepage"`
	Keywords    []string          `json:"keywords"`
	Scripts     PackageScripts    `json:"scripts"`
	Repository  PackageRepository `json:"repository"`
}

func UnmarshalNPMFile(filepath string) Package {
	var pkg Package
	var jsonData, _ = commons.ReadFile(filepath)
	json.Unmarshal([]byte(jsonData), &pkg)
	return pkg
}

func MarshalNPMFile(pkg Package, filepath string) {
	result, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		panic(err)
	}

	commons.ToFile(filepath, result)
}
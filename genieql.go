// Package genieql generates code to interact with a database.
package genieql

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Build Tag constants
const (
	BuildTagIgnore   = "genieql.ignore"   // used to filter out files from the build context using !genieql.ignore
	BuildTagGenerate = "genieql.generate" // used to specify which files should be analyzed for directives using genieql.generate
)

// Preface text inserted at the top of all generated files.
const Preface = `

// DO NOT EDIT: This File was auto generated by the following command:
// genieql %s
`

// ConfigurationDirectory determines the configuration directory based on the
// go environment.
func ConfigurationDirectory() string {
	var (
		err         error
		defaultPath string
	)

	if defaultPath, err = FindModuleRoot("."); err == nil && defaultPath != "" {
		return filepath.Join(defaultPath, ".genieql")
	}

	paths := filepath.SplitList(os.Getenv("GOPATH"))

	if len(paths) == 0 {
		if defaultPath, err = os.Getwd(); err != nil {
			log.Fatalln(err)
		}
	} else {
		defaultPath = paths[0]
	}

	return filepath.Join(defaultPath, ".genieql")
}

// PrintColumnInfo ...
func PrintColumnInfo(columns ...ColumnInfo) {
	for _, column := range columns {
		log.Println(column)
	}
}

// FindModuleRoot pulled from: https://github.com/golang/go/blob/src/cmd/dist/build.go#L1595
func FindModuleRoot(dir string) (cleaned string, err error) {
	if dir == "" {
		return "", errors.New("cannot located go.mod from a blank directory path")
	}

	if cleaned, err = filepath.Abs(filepath.Clean(dir)); err != nil {
		return "", errors.Wrap(err, "failed to determined absolute path to directory")
	}

	// Look for enclosing go.mod.
	for {
		gomod := filepath.Join(cleaned, "go.mod")
		if fi, err := os.Stat(gomod); err == nil && !fi.IsDir() {
			return cleaned, nil
		}

		d := filepath.Dir(cleaned)

		if d == cleaned {
			break
		}

		cleaned = d
	}

	return "", errors.Errorf("go.mod not found: %s", dir)
}

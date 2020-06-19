// Package genieql generates code to interact with a database.
package genieql

import (
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/tools/imports"

	"bitbucket.org/jatone/genieql/internal/iox"
)

// Build Tag constants
const (
	BuildTagIgnore   = "genieql.ignore"   // used to filter out files from the build context using !genieql.ignore
	BuildTagGenerate = "genieql.generate" // used to specify which files should be analyzed for directives using genieql.generate
)

// Preface text inserted at the top of all generated files.
const Preface = `

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql %s
`

// FormatOutput formats and resolves imports for the raw bytes representing a go
// source file and writes them into the dst.
func FormatOutput(dst io.Writer, raw []byte) (err error) {
	if raw, err = imports.Process("generated.go", raw, nil); err != nil {
		return errors.Wrap(err, "failed to add required imports")
	}

	if raw, err = format.Source(raw); err != nil {
		return errors.Wrap(err, "failed to format source")
	}

	_, err = dst.Write(raw)
	return errors.Wrap(err, "failed to write to completed code to destination")
}

// Reformat a file
func Reformat(in io.ReadWriteSeeker) (err error) {
	var (
		raw []byte
	)

	// ensure we're at the start of the file.
	if err = iox.Rewind(in); err != nil {
		return err
	}

	if raw, err = ioutil.ReadAll(in); err != nil {
		return err
	}

	if raw, err = imports.Process("generated.go", []byte(string(raw)), nil); err != nil {
		return errors.Wrap(err, "failed to add required imports")
	}

	// ensure we're at the start of the file.
	if err = iox.Rewind(in); err != nil {
		return err
	}

	if _, err = in.Write(raw); err != nil {
		return errors.Wrap(err, "failed to write formatted content")
	}

	return nil
}

// ReformatFile a file
func ReformatFile(in *os.File) (err error) {
	var (
		raw []byte
	)

	// ensure we're at the start of the file.
	if err = iox.Rewind(in); err != nil {
		return err
	}

	if raw, err = ioutil.ReadAll(in); err != nil {
		return err
	}

	if raw, err = imports.Process("generated.go", []byte(string(raw)), nil); err != nil {
		return errors.Wrap(err, "failed to add required imports")
	}

	// ensure we're at the start of the file.
	if err = iox.Rewind(in); err != nil {
		return err
	}

	if err = in.Truncate(0); err != nil {
		return errors.Wrap(err, "failed to truncate file")
	}

	if _, err = in.Write(raw); err != nil {
		return errors.Wrap(err, "failed to write formatted content")
	}

	return nil
}

// Format arbitrary source fragment.
func Format(s string) (_ string, err error) {
	var (
		raw []byte
	)

	if raw, err = imports.Process("generated.go", []byte(s), &imports.Options{Fragment: true, Comments: true, TabIndent: true, TabWidth: 8}); err != nil {
		return "", errors.Wrap(err, "failed to add required imports")
	}

	if raw, err = format.Source(raw); err != nil {
		return "", errors.Wrap(err, "failed to format source")
	}

	return string(raw), nil
}

// LoadInformation loads table information based on the configuration and
// table name.
func LoadInformation(configuration Configuration, table string) (details TableDetails, err error) {
	var (
		driver  Driver
		dialect Dialect
	)

	if dialect, err = LookupDialect(configuration); err != nil {
		return details, err
	}

	if driver, err = LookupDriver(configuration.Driver); err != nil {
		return details, err
	}

	details, err = LookupTableDetails(driver, dialect, table)

	return details, err
}

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

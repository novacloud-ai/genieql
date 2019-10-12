package cmd

import (
	"io"
	"log"
	"os"

	"github.com/pkg/errors"

	"bitbucket.org/jatone/genieql"
)

type errWriter struct {
	err error
}

func (t errWriter) Write([]byte) (int, error) {
	return 0, t.err
}

func (t errWriter) Close() error {
	return t.err
}

// DefaultWriteFlags default write flags for WriteStdoutOrFile
const DefaultWriteFlags = os.O_CREATE | os.O_TRUNC | os.O_RDWR

// WriteStdoutOrFile writes to stdout if fpath is empty.
func WriteStdoutOrFile(g genieql.Generator, fpath string, flags int) error {
	var (
		err error
		dst io.WriteCloser = os.Stdout
	)

	if len(fpath) > 0 {
		log.Println("writing results to", fpath)
		if dst, err = os.OpenFile(fpath, flags, 0666); err != nil {
			dst = errWriter{err: errors.Wrap(err, "")}
		}
		defer dst.Close()
	}

	if err = g.Generate(dst); err != nil {
		log.Printf("%s: failed to generate: %+v\n", genieql.PrintDebug(), err)
		return err
	}

	return err
}
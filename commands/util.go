package commands

import (
	"bytes"
	"io"
	"log"
	"os"

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
		err    error
		buffer                = bytes.NewBuffer([]byte{})
		dst    io.WriteCloser = os.Stdout
	)

	if err = g.Generate(buffer); err != nil {
		return err
	}

	if len(fpath) > 0 {
		log.Println("Writing Results to", fpath)
		if dst, err = os.OpenFile(fpath, flags, 0666); err != nil {
			dst = errWriter{err: err}
		}
		defer dst.Close()
	}

	_, err = io.Copy(dst, buffer)
	return err
}

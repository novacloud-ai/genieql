package generators_test

import (
	"io"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGenerators(t *testing.T) {
	log.SetOutput(io.Discard)
	log.SetFlags(log.Flags() | log.Lshortfile)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generators Suite")
}

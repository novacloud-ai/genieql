package generators_test

import (
	"bytes"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"

	"bitbucket.org/jatone/genieql"
	_ "bitbucket.org/jatone/genieql/internal/drivers"

	. "bitbucket.org/jatone/genieql/generators"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scanner", func() {
	config := genieql.MustConfiguration(
		genieql.ConfigurationOptionLocation(
			filepath.Join(genieql.ConfigurationDirectory(), "scanner-test.config"),
		),
	)
	genieql.RegisterDriver(config.Driver, noopDriver{})

	DescribeTable("should build a scanner for builtin types",
		func(definition, fixture string) {
			buffer := bytes.NewBuffer([]byte{})
			formatted := bytes.NewBuffer([]byte{})
			fset := token.NewFileSet()

			node, err := parser.ParseFile(fset, "example", definition, 0)
			Expect(err).ToNot(HaveOccurred())

			buffer.WriteString("package example\n\n")
			for _, d := range genieql.SelectFuncType(genieql.FindTypes(node)...) {
				for _, g := range ScannerFromGenDecl(d, ScannerOptionConfiguration(config)) {
					Expect(g.Generate(buffer)).ToNot(HaveOccurred())
					buffer.WriteString("\n")
				}
			}
			expected, err := ioutil.ReadFile(fixture)
			Expect(err).ToNot(HaveOccurred())
			Expect(genieql.FormatOutput(formatted, buffer.Bytes())).ToNot(HaveOccurred())
			Expect(formatted.String()).To(Equal(string(expected)))
		},
		Entry("scanner int", `package example; type ExampleInt func(arg int)`, ".fixtures/scanners/int.go"),
		Entry("scanner bool", `package example; type ExampleBool func(arg bool)`, ".fixtures/scanners/bool.go"),
		Entry("scanner time.Time", `package example; type ExampleTime func(arg time.Time)`, ".fixtures/scanners/time.go"),
		Entry("scanner multipleParams", `package example; type ExampleMultipleParam func(arg1, arg2 int, arg3 bool, arg4 string)`, ".fixtures/scanners/multiple_params.go"),
	)

	DescribeTable("should build scanners with only the specified outputs",
		func(definition, fixture string, options ...ScannerOption) {
			buffer := bytes.NewBuffer([]byte{})
			formatted := bytes.NewBuffer([]byte{})
			fset := token.NewFileSet()

			node, err := parser.ParseFile(fset, "example", definition, 0)
			Expect(err).ToNot(HaveOccurred())

			buffer.WriteString("package example\n\n")
			for _, d := range genieql.SelectFuncType(genieql.FindTypes(node)...) {
				for _, g := range ScannerFromGenDecl(d, append(options, ScannerOptionConfiguration(config))...) {
					Expect(g.Generate(buffer)).ToNot(HaveOccurred())
					buffer.WriteString("\n")
				}
			}
			expected, err := ioutil.ReadFile(fixture)
			Expect(err).ToNot(HaveOccurred())
			Expect(genieql.FormatOutput(formatted, buffer.Bytes())).ToNot(HaveOccurred())
			Expect(formatted.Bytes()).To(Equal(expected))
		},
		Entry("scanner int without interface",
			`package example; type ExampleIntNoInterface func(arg int)`,
			".fixtures/scanners/int_without_interface.go",
			ScannerOptionOutputMode(ModeStatic|ModeDynamic),
		),
		Entry("scanner int without static",
			`package example; type ExampleIntNoStatic func(arg int)`,
			".fixtures/scanners/int_without_static.go",
			ScannerOptionOutputMode(ModeInterface|ModeDynamic),
		),
		Entry("scanner int without dynamic",
			`package example; type ExampleIntNoDynamic func(arg int)`,
			".fixtures/scanners/int_without_dynamic.go",
			ScannerOptionOutputMode(ModeInterface|ModeStatic),
		),
	)
})

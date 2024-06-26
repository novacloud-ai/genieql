package dialects

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dialect", func() {
	Describe("dialectRegistry", func() {
		Describe("RegisterDialect", func() {
			It("should err if the dialect is already registered", func() {
				dialect := TestFactory{}
				reg := dialectRegistry{}
				Expect(reg.RegisterDialect("testDialect", dialect)).ToNot(HaveOccurred())
				Expect(reg.RegisterDialect("testDialect", dialect)).To(MatchError(ErrDuplicateDialect))
			})

			It("should register a dialect", func() {
				dialect := TestFactory{}
				reg := dialectRegistry{}
				Expect(reg.RegisterDialect("testDialect", dialect)).ToNot(HaveOccurred())
			})
		})

		Describe("LookupDialect", func() {
			It("should err if the dialect is not registered", func() {
				reg := dialectRegistry{}
				dialect, err := reg.LookupDialect("testDialect")
				Expect(dialect).To(BeNil())
				Expect(err).To(MatchError("dialect (testDialect) is not registered"))
			})

			It("should return the dialect if its been registered", func() {
				dialectName := "testDialect"
				dialect := TestFactory{}
				reg := dialectRegistry{}
				Expect(reg.RegisterDialect(dialectName, dialect)).ToNot(HaveOccurred())
				foundDialect, err := reg.LookupDialect(dialectName)
				Expect(err).ToNot(HaveOccurred())
				Expect(foundDialect).To(Equal(dialect))
			})
		})
	})
})

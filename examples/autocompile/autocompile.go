package autocompile

//go:generate genieql auto -o "genieql.gen.go"
// to invoke: go generate ./examples/autocompile/...
// export GOBIN=${HOME}/go/bin
// go install bitbucket.org/jatone/genieql/cmd/...
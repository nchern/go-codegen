[![Go Report Card](https://goreportcard.com/badge/github.com/nchern/go-codegen)](https://goreportcard.com/report/github.com/nchern/go-codegen)

# go-codegen

With a lack of templates some of the gophers really miss them...
Yet another attempt of code generaton in Golang.

## Install
```go
go get -u "github.com/nchern/go-codegen/..."
```

## Testing
```bash
make test # generates code and runs tests
```

## Usage
```text
Go code generation tool. Prints output to stdout

Usage:
  go-codegen [command]

Available Commands:
  generic     Processes go source as generic file and outputs code with substituted type vars
  help        Help about any command
  immutable   Generates immutable implementation by a given interface.

Flags:
  -f, --file string   input file name (reqiured)
  -h, --help          help for go-codegen
  -p, --pkg string    Golang package name. Substitutes existing package name or makes generator to add one

Use "go-codegen [command] --help" for more information about a command.
```

### Generics

```bash
# Outputs typed string list implementation to stdout
go-codegen --pkg=main generic -f pkg/generic/list/list.go string
```

Generics are implemented by parsing an input go source file into an AST and substituting predefined "generic" types. Currently only a fixed list of such types is supported: `T0, T1, ..., T5`. The advantage of such approach as the generic implementation(input source) is the correct go source that can be tested.
See build-ins as [examples](pkg/generic/list/list.go) of how to define a generic type.
See [test code](tests/generic) for more examples.

### Immutables
An experiment. Inspired by Java Immutables. See more examples in [test code](tests/immutable/)

## Nice to have

 * Integration with [gen](http://alikewise.com/gen/)


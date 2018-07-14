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
```bash
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
  -p, --pkg string    Golang package name. Substitues existing package name or makes generator to add one

Use "go-codegen [command] --help" for more information about a command.
```

### Generics

```bash
# Outputs typed string list implementation to stdout
go-codegen --pkg=main generic -f pkg/generic/list/list.go string
```

See build-ins as [examples](pkg/generic/list/list.go) of how to define a generic type.
See [test code](tests) for more examples.

### Immutables
TBD.
 - [ ] implement JSON serialisation / deserialisation
 - [ ] examples

# Nice to have

 * Integration with [gen](http://alikewise.com/gen/)


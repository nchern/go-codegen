[![Go Report Card](https://goreportcard.com/badge/github.com/nchern/go-codegen)](https://goreportcard.com/report/github.com/nchern/go-codegen)
[![Coverage](https://gocover.io/_badge/github.com/nchern/go-codegen)](https://gocover.io/github.com/nchern/go-codegen)


# go-codegen

With a lack of templates some of the gophers really miss them...
Yet another attempt of code generaton in Golang.

## Testing
```bash
make test # generates code and runs tests
```

## Install
```bash
make install
```

## Usage
```bash
# Prints help
$ go-codegen -h
```

### Generics

#### Usage
```bash
# Outputs built-in string to int map implementation to stdout
go-codegen generic -t=typedmap string int

# Reads generic source from file pkg/generic/list/list.go and outputs typed implementation to stdout
go-codegen generic -f pkg/generic/list/list.go string

# More help on available built-ins and other command line params
go-codegen help generic
```

Generics are implemented by parsing an input go source file into an AST and substituting predefined "generic" types. Currently only a fixed list of such types is supported: `T0, T1, ..., T5`. The advantage of such approach as the generic implementation(input source) is the correct go source that can be tested.
See build-ins as [examples](pkg/generic/list/list.go) of how to define a generic type.
See [test code](tests/generic) for more examples.

### Immutables

#### Usage
```bash
# Gets a .go source file with defined interfaces and outputs implementation along with builder class to create instances
go-codegen --pkg=model immutable -f tests/immutable/model/model.go
```

An experiment. Inspired by Java Immutables. See more examples in [test code](tests/immutable/).

### Constructor

Generates constructor function for struct. More details see [here](docs/constructor.md).

### Impl

Generates interface implementations stubs. More details see [here](docs/impl.md).

## Nice to have

 * Integration with [gen](http://alikewise.com/gen/)

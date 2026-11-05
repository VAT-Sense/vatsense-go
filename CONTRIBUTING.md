## Setting up the environment

To set up the repository, run:

```sh
$ ./scripts/bootstrap
$ ./scripts/lint
```

This will install all the required dependencies and build the SDK.

You can also [install go 1.22+ manually](https://go.dev/doc/install).

## Modifying/Adding code

The SDK was first generated from the OpenAPI spec by Stainless and is now maintained by hand.
When the API changes, update the types and methods to match `https://vatsense.com/openapi.yaml`,
keeping the existing names so code that uses the SDK keeps working. Header comments in the files
still say they were generated; they are kept to keep the diff small.

## Adding and running examples

All files in the `examples/` directory are not modified by the generator and can be freely edited or added to.

```go
# add an example to examples/<your-example>/main.go

package main

func main() {
  // ...
}
```

```sh
$ go run ./examples/<your-example>
```

## Using the repository from source

To use a local version of this library from source in another project, edit the `go.mod` with a replace
directive. This can be done through the CLI with the following:

```sh
$ go mod edit -replace github.com/VAT-Sense/vatsense-go=/path/to/vatsense-go
```

## Running tests

```sh
$ ./scripts/test
```

## Formatting

This library uses the standard gofmt code formatter:

```sh
$ ./scripts/format
```

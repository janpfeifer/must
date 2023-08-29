# "Must" Shortcuts

[![GoDev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/janpfeifer/must?tab=doc)
[![GitHub](https://img.shields.io/github/license/janpfeifer/must)](https://github.com/Kwynto/gosession/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/janpfeifer/must)](https://goreportcard.com/report/github.com/janpfeifer/must)

A trivial package to provide a collection of functions shortcuts error returns to panics.

This is handy when doing small programs, or script-like programs in Go, where any failure should just panic.

It defines:

```go
func M(err error) {
  if err != nil {
    panic(err)
  }
}

func M1[T any](value T, err error) T {
  M(err)
  return value
}

// And similarly it defines functions M2 ... to ... M9.
```

Example of usage:

```go
package main

import . "github.com/janpfeifer/must"

func main() {
  fName := "my_file.txt"
  contents := M1(os.ReadFile(fName))
  M(os.Remove(fName))
}
```

Or, if using a Go Jupyter Notebook with [GoNB](https://github.com/janpfeifer/gonb):

```go
import . "github.com/janpfeifer/must"
%%
fName := "my_file.txt"
contents := M1(os.ReadFile(fName))
M(os.Remove(fName))
```

Notice you can also redefine `M`, and all other functions will automatically follow suit:

```go
import . "github.com/janpfeifer/must"

func init() {
	M = func(err error) {
	    if err != nil {
		    log.Fatalf("Interrupted with error: %+v", err)	
        }       	
    }
}

%%
contents := M1(os.ReadFile("my_file"))  // This will fail with log.Fatalf if err != nil.
```

> [!NOTE]
> `panic` is like an exception, and it can be caught (with `recover`), if you need it at some point in the "script".

## FAQ

1. What if I don't want to `panic` but do something else ? <br> See example of reassigning `M` to however you want errors to be handled.
2. What if I don't like `M` as the name for these functions, and want something else ? <br> If it's something generic, just create an issue, I'll quickly create a subpackage with your favourite naming convention.


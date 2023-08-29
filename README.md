# "Must" Shortcuts

A trivial package to provide a collection of functions shortcuts error returns to panics.

This is handy when doing small programs, or script-like programs in Go, where any failure should just panic.

It defines:

```go
func M(err error) {
  if err != nil {
    panice(err)
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

```
import . "github.com/janpfeifer/must"
%%
fName := "my_file.txt"
contents := M1(os.ReadFile(fName))
M(os.Remove(fName))
```

**Note**: `panic` is like an exception, and it can be caught (with `recover`), if you need it at some point in the "script".

> [!NOTE]
> Highlights information that users should take into account, even when skimming.


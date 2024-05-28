# See the [got-reload README](https://github.com/got-reload/got-reload/blob/main/README.md) for what this is
and what it's for.

# Original code
```go
var (
	f = new(float64)
)

func F1() int {
	*f += 0.1
	fmt.Printf("f: %0.3f, sin %0.3f\n", *f, example2.Sin(*f))
	return 1
}
```

# Filtered (reloadable) code
```go
var (
	GRLx_f = new(float64)
)

func F1() int { return GRLfvar_F1() }

var GRLfvar_F1 func() int

func init() {
	GRLfvar_F1 = func() int {
		*GRLx_f += 0.1
		fmt.Printf("f: %0.3f, sin %0.3f\n", *GRLx_f, example2.Sin(*GRLx_f))
		return 1
	}
}
```

# Example reload code
Note changed bit at `*GRLx_f += 0.2`.

```go
package main

import (
       fmt "fmt"

       . "github.com/got-reload/demo/example"
       example2 "github.com/got-reload/demo/example2"
)

func main() {
      GRLfvar_F1 = func() int {
              *GRLx_f += 0.2 // note change here
              fmt.Printf("f: %0.3f, sin %0.3f\n", *GRLx_f, example2.Sin(*GRLx_f))
              return 1
      }
}
```

# Sample run
(See `# NOTE: ...` annotations.)

```console
main.go:173: copying [...]/github.com/got-reload/demo to /tmp/got-reload
main.go:303: Parsing package [github.com/got-reload/demo/example]
main.go:195: GOT_RELOAD_PKGS=github.com/got-reload/demo/example
main.go:195: GOT_RELOAD_START_RELOADER=1
main.go:195: GOT_RELOAD_SOURCE_DIR=[...]/github.com/got-reload/demo
github.com/got-reload/demo/example
github.com/got-reload/demo
GRL: 11:24:22.871087 reloader.go:114: Running go list from [...]/github.com/got-reload/demo
GRL: 11:24:22.921793 reloader.go:196: Starting reloader
GRL: 11:24:23.436872 reloader.go:154: WatchedPkgs: [github.com/got-reload/demo/example], PkgsToDirs: map[github.com/got-reload/demo/example:[...]/github.com/got-reload/demo/example], DirsToPkgs: map[[...]/github.com/got-reload/demo/example:github.com/got-reload/demo/example]
GRL: 11:24:23.436971 reloader.go:161: Watching [...]/github.com/got-reload/demo/example
Press enter to call example.F1 and example2.F2 repeatedly
Enter s to stop
example2.Sin: f: 0.100
f: 0.100, sin 0.100    # NOTE: f starts at 0.1
example.F1: 1
GRL: 11:24:24.438943 reloader.go:215: Reloader waiting for all RegisterAll calls to finish
GRL: 11:24:24.439034 reloader.go:220: Reloader continuing

example2.Sin: f: 0.200
f: 0.200, sin 0.199    # NOTE: f increases by 0.1
example.F1: 1

example2.Sin: f: 0.300
f: 0.300, sin 0.296
example.F1: 1
GRL: 11:24:32.364846 reloader.go:241: Changed: [...]/github.com/got-reload/demo/example/example.go
GRL: 11:24:32.493786 reloader.go:265: Reparsing package containing [...]/github.com/got-reload/demo/example/example.go
GRL: 11:24:32.913444 reloader.go:276: Refiltering package containing [...]/github.com/got-reload/demo/example/example.go
GRL: 11:24:32.913503 reloader.go:305: Looking for updated functions in github.com/got-reload/demo/example
GRL: 11:24:32.913642 reloader.go:360: GRLfvar_F1 is new
GRL: 11:24:32.916578 reloader.go:438: Ran GRLfvar_F1

example2.Sin: f: 0.500
f: 0.500, sin 0.479    # NOTE: f increases by 0.2
example.F1: 1

example2.Sin: f: 0.700
f: 0.700, sin 0.644
example.F1: 1
s
```

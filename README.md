# stacktrace
Debug code by stacktrace

## Install:
```sh
go get -u github.com/bookerzzz/stacktrace
```

## Usage:
Example is included in `stacktrace_test.go`.

```go
import "github.com/bookerzzz/stacktrace"

// quick and easy
stacktrace.Print(os.Stdout)

// or if you want colourful output
stacktrace.Pretty(os.Stdout)

// or for customised output
info := stacktrace.Get()
for i := len(info) - 1; i > -1; i-- {
    v := info[i]
    fmt.Printf("%02d: %s in %s:%d\n", i, v.FunctionName, v.FileName, v.FileLine)
}

// Output
// 09: goexit in /usr/local/go/src/runtime/asm_amd64.s:2337
// 08: main in /usr/local/go/src/runtime/proc.go:185
// 07: main in ./stacktrace/_test/_testmain.go:44
// 06: Run in /usr/local/go/src/testing/testing.go:922
// 05: runExamples in /usr/local/go/src/testing/example.go:46
// 04: runExample in /usr/local/go/src/testing/example.go:122
// 03: ExamplePrint in ./stacktrace/stacktrace_test.go:21
// 02: boo in ./stacktrace/stacktrace_test.go:9
// 01: foo in ./stacktrace/stacktrace_test.go:13
// 00: display in ./stacktrace/stacktrace_test.go:17
```

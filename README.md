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

stacktrace.Print(os.Stdout, "your_package_name")

// or for customised output

info := stacktrace.Get("your_package_name")
for i := len(info) - 1; i > -1; i-- {
    v := info[i]
    fmt.Printf("%02d: [Function]%s [File]%s:%d\n", i, v.FunctionName, v.FileName, v.FileLine)
}
// Output
// 09: [Function]goexit [File]/usr/local/go/src/runtime/asm_amd64.s:2337
// 08: [Function]main [File]/usr/local/go/src/runtime/proc.go:185
// 07: [Function]main [File]./stacktrace/_test/_testmain.go:44
// 06: [Function]Run [File]/usr/local/go/src/testing/testing.go:922
// 05: [Function]runExamples [File]/usr/local/go/src/testing/example.go:46
// 04: [Function]runExample [File]/usr/local/go/src/testing/example.go:122
// 03: [Function]ExamplePrint [File]./stacktrace/stacktrace_test.go:21
// 02: [Function]boo [File]./stacktrace/stacktrace_test.go:9
// 01: [Function]foo [File]./stacktrace/stacktrace_test.go:13
// 00: [Function]display [File]./stacktrace/stacktrace_test.go:17
```

# stacktrace
Debug code by stacktrace

## Install:
```sh
go get github.com/bookerzzz/stacktrace
```

## Usage:
Example is included in `stacktrace_test.go`.

```go
import "github.com/bookerzzz/stacktrace"

stacktrace.OutputStackTrace(os.Stdout, "your_package_name")

// or for customised output

info := stacktrace.GetStackTrace("your_package_name")
for i := len(info) - 1; i > -1; i-- {
    v := info[i]
    fmt.Printf("%02d: [Function]%s [File]%s:%d\n", i, v.FunctionName, v.FileName, v.FileLine)
}
// Output
// 03: [Function]goexit [File]/usr/local/Cellar/go/1.9.2/libexec/src/runtime/asm_amd64.s:2337
// 02: [Function]tRunner [File]/usr/local/Cellar/go/1.9.2/libexec/src/testing/testing.go:746
// 01: [Function]TestGetStackTrace [File]./stacktrace/stacktrace_test.go:35
// 00: [Function]GetStackTrace [File]./stacktrace/stacktrace.go:57}
```

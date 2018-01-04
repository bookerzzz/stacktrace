package stacktrace

import (
	"fmt"
	"os"
)

func boo() {
	foo()
}

func foo() {
	display()
}

func display() {
	Print(os.Stdout, "bookerzzz")
}

func ExamplePrint() {
	boo()
	// Output
	// 06: [Function]goexit [File]/usr/local/Cellar/go/1.9.2/libexec/src/runtime/asm_amd64.s:2337
	// 05: [Function]tRunner [File]/usr/local/Cellar/go/1.9.2/libexec/src/testing/testing.go:746
	// 04: [Function]TestOutputStackTrace [File]./stacktrace/stacktrace_test.go:23
	// 03: [Function]boo [File]./stacktrace/stacktrace_test.go:10
	// 02: [Function]foo [File]./stacktrace/stacktrace_test.go:14
	// 01: [Function]display [File]./stacktrace/stacktrace_test.go:18
	// 00: [Function]Print [File]./stacktrace/stacktrace.go:62}
}

func ExampleGet() {
	info := Get("bookerzzz")
	for i := len(info) - 1; i > -1; i-- {
		v := info[i]
		fmt.Printf("%02d: [Function]%s [File]%s:%d\n", i, v.FunctionName, v.FileName, v.FileLine)
	}
	// Output
	// 03: [Function]goexit [File]/usr/local/Cellar/go/1.9.2/libexec/src/runtime/asm_amd64.s:2337
	// 02: [Function]tRunner [File]/usr/local/Cellar/go/1.9.2/libexec/src/testing/testing.go:746
	// 01: [Function]TestGetStackTrace [File]./stacktrace/stacktrace_test.go:35
	// 00: [Function]Get [File]./stacktrace/stacktrace.go:57}
}

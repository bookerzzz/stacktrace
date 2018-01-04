package stacktrace

import (
	"fmt"
	"io"
	"regexp"
	"runtime"
	"strings"
)

var (
	re = regexp.MustCompile(`^(\S.+)\.(\S.+)$`)
)

// CallerInfo is stack trace information including package name, function name, file name, file line
type CallerInfo struct {
	PackageName  string
	FunctionName string
	FileName     string
	FileLine     int
}

func formatPath(path, separator string) string {
	ret := strings.Split(path, separator)
	if len(ret) > 1 {
		return fmt.Sprintf(".%s", ret[len(ret)-1])
	}
	return path
}

func dump(separator string) (callerInfo []*CallerInfo) {
	for i := 1; ; i++ {
		pc, _, _, ok := runtime.Caller(i) // https://golang.org/pkg/runtime/#Caller
		if !ok {
			break
		}

		fn := runtime.FuncForPC(pc)
		fileName, fileLine := fn.FileLine(pc)

		// format path
		if separator != "" {
			fileName = formatPath(fileName, separator)
		}

		additionalInfo := re.FindStringSubmatch(fn.Name())
		callerInfo = append(callerInfo, &CallerInfo{
			PackageName:  additionalInfo[1],
			FunctionName: additionalInfo[2],
			FileName:     fileName,
			FileLine:     fileLine,
		})
	}
	// Remove first element from callerInfo because it's a stacktrace internal function
	return callerInfo[1:]
}

// Get slice of CallerInfo addresses
func Get(separator string) []*CallerInfo {
	info := dump(separator)
	return info
}

// Print stack trace information to given io.Writer
func Print(w io.Writer, separator string) {
	info := dump(separator)
	for i := len(info) - 1; i > -1; i-- {
		v := info[i]
		fmt.Fprintf(w, "%02d: [Function]%s [File]%s:%d\n", i, v.FunctionName, v.FileName, v.FileLine)
	}
}

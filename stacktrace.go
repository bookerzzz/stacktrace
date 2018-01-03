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

func dumpStackTrace(separator string) (callerInfo []*CallerInfo) {
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
	return
}

// GetStackTrace returns slice of CallerInfo address
func GetStackTrace(separator string) []*CallerInfo {
	info := dumpStackTrace(separator)
	return info
}

// OutputStackTrace output stack trace information by io.Writer
func OutputStackTrace(w io.Writer, separator string) {
	info := dumpStackTrace(separator)
	for i := len(info) - 1; i > -1; i-- {
		v := info[i]
		fmt.Fprintf(w, "%02d: [Function]%s [File]%s:%d\n", i, v.FunctionName, v.FileName, v.FileLine)
	}
}

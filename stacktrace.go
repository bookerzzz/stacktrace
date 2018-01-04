package stacktrace

import (
	"fmt"
	"io"
	"regexp"
	"runtime"
	"strings"
	"os"
)

type Colour string

var (
	re = regexp.MustCompile(`^(\S.+)\.(\S.+)$`)
	colourReset  Colour = "\x1B[0m"
	colourRed    Colour = "\x1B[38;5;124m"
	colourYellow Colour = "\x1B[38;5;208m"
	colourBlue   Colour = "\x1B[38;5;33m"
	colourGrey   Colour = "\x1B[38;5;144m"
	colourGreen  Colour = "\x1B[38;5;34m"
	colourGold   Colour = "\x1B[38;5;3m"
)

// CallerInfo is stack trace information including package name, function name, file name, file line
type CallerInfo struct {
	PackageName  string
	FunctionName string
	FileName     string
	FileLine     int
}

func colourf(format string, c Colour, a ...interface{}) string {
	return colour(fmt.Sprintf(format,  a...), c)
}

func colour(str string, c Colour) string {
	return string(c)+str+string(colourReset)
}

func formatPath(path, prefix string) string {
	if strings.HasPrefix(path, prefix) {
		return fmt.Sprintf(".%s", strings.TrimPrefix(path, prefix))
	}
	return path
}

func dump() (callerInfo []*CallerInfo) {
	for i := 1; ; i++ {
		// https://golang.org/pkg/runtime/#Caller
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			break
		}

		fn := runtime.FuncForPC(pc)
		fileName, fileLine := fn.FileLine(pc)

		// format path
		fileName = formatPath(fileName, os.Getenv("PWD"))

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
func Get() []*CallerInfo {
	info := dump()
	return info
}

// Print stack trace information to given io.Writer
func Print(w io.Writer) {
	info := dump()
	for i := len(info) - 1; i > -1; i-- {
		v := info[i]
		fmt.Fprintf(w, "%02d: [Function]%s [File]%s:%d\n", i, v.FunctionName, v.FileName, v.FileLine)
	}
}

// Pretty print stack trace information to given io.Writer
func Pretty(w io.Writer) {
	info := dump()
	for i := len(info) - 1; i > -1; i-- {
		v := info[i]
		fmt.Fprintf(w, "%s %s %s\n", colourf("%02d", colourYellow, i), colour(v.FunctionName, colourGreen), colourf("%s:%d", colourBlue, v.FileName, v.FileLine))
	}
}

// based on source from https://github.com/gin-gonic/gin/blob/master/mode.go

package gojenkins

import (
	"fmt"
	"strings"
)

// IsDebugging returns true if the framework is running in debug mode.
// Use SetMode(gojenkins.ReleaseMode) to disable debug mode.
func IsDebugging() bool {
	return jenkinsMode == debugCode
}

func debugPrint(format string, values ...interface{}) {
	if IsDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(DefaultWriter, "[GOJENKINS-debug] "+format, values...)
	}
}

func debugPrintError(err error) {
	if err != nil {
		if IsDebugging() {
			fmt.Fprintf(DefaultErrorWriter, "[GOJENKINS-debug] [ERROR] %v\n", err)
		}
	}
}

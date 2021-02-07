// based on source from https://github.com/gin-gonic/gin/blob/master/mode.go

package gojenkins

import (
	"io"
	"os"
)

// EnvJenkinsMode indicates environment name for jenkins mode.
const EnvJenkinsMode = "JENKINS_MODE"

const (
	// DebugMode indicates jenkins mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates jenkins mode is release.
	ReleaseMode = "release"
	// TestMode indicates jenkins mode is test.
	TestMode = "test"
)
const (
	debugCode = iota
	releaseCode
	testCode
)

// DefaultWriter is the default io.Writer used by Jenkins for debug output and
// middleware output like Logger() or Recovery().
// Note that both Logger and Recovery provides custom ways to configure their
// output io.Writer.
// To support coloring in Windows use:
// 		import "github.com/mattn/go-colorable"
// 		gin.DefaultWriter = colorable.NewColorableStdout()
var DefaultWriter io.Writer = os.Stdout

// DefaultErrorWriter is the default io.Writer used by Jenkins to debug errors
var DefaultErrorWriter io.Writer = os.Stderr

var jenkinsMode = debugCode
var modeName = DebugMode

func init() {
	mode := os.Getenv(EnvJenkinsMode)
	SetMode(mode)
}

// SetMode sets jenkins mode according to input string.
func SetMode(value string) {
	switch value {
	case DebugMode, "":
		jenkinsMode = debugCode
	case ReleaseMode:
		jenkinsMode = releaseCode
	case TestMode:
		jenkinsMode = testCode
	default:
		panic("jenkins mode unknown: " + value)
	}
	if value == "" {
		value = DebugMode
	}
	modeName = value
}

// Mode returns currently jenkins mode.
func Mode() string {
	return modeName
}

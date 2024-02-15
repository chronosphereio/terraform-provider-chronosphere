package buildinfo

import (
	"log"
	"os"
	"runtime"
)

// All of these vars are replaced at link time using ldflags.
var (
	// Version is a valid semantic version
	Version = "unknown"
	// SHA contains the git SHA of the commit
	SHA = "unknown"
	// Date is the date of the build
	Date = "unknown"
	// LogBuildInfoAtStartup controls whether we log build information at startup. If its
	// set to a non-empty string, we log the build information at process startup.
	LogBuildInfoAtStartup string
	// LogBuildInfoToStdout controls whether we log build information to stdout or stderr.
	// If it is set to a non-empty string then the build info will be logged to stdout,
	// otherwise it will be logged to stderr (assuming LogBuildInfoAtStartup is also
	// non-empty).
	LogBuildInfoToStdout string
)

// LogBuildInfo logs the build information using the default logger.
func LogBuildInfo() {
	LogBuildInfoWithLogger(log.Default())
}

// LogBuildInfoWithLogger logs the build information using the provided logger.
func LogBuildInfoWithLogger(logger *log.Logger) {
	logger.Printf("Go Runtime version: %s\n", runtime.Version())
	logger.Printf("Build Version:      %s\n", Version)
	logger.Printf("Build SHA:     %s\n", SHA)
	logger.Printf("Build Date:         %s\n", Date)
}

func init() {
	if LogBuildInfoAtStartup != "" {
		logger := log.Default()
		if LogBuildInfoToStdout != "" {
			logger = log.New(os.Stdout, "", log.LstdFlags)
		}
		LogBuildInfoWithLogger(logger)
	}
}

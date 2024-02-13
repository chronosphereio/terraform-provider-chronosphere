// Package buildvar is used for version information set at build time
// using ldflags -X.
package buildvar

import (
	"log"
	"os"
	"runtime"
)

var (
	// Revision is the VCS revision associated with this build. Overridden using ldflags
	// at compile time.
	Revision = "unknown"

	// Branch is the VCS branch associated with this build.
	Branch = "unknown"

	// Version is the version associated with this build.
	Version = "unknown"

	// BuildDate is the date this build was created.
	BuildDate = "unknown"

	// BuildTimeUnix is the seconds since epoch representing the date this build was created.
	BuildTimeUnix = "0"

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
	logger.Printf("Build Revision:     %s\n", Revision)
	logger.Printf("Build Branch:       %s\n", Branch)
	logger.Printf("Build Date:         %s\n", BuildDate)
	logger.Printf("Build TimeUnix:     %s\n", BuildTimeUnix)
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

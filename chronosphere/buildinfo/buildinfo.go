// Copyright 2023 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

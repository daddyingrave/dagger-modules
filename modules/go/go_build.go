// Go programming language module.
package main

import (
	"strings"

	"dagger/go/internal/dagger"
)

// Build Compile the packages into a binary.
func (m *Go) Build(
	// Package to compile.
	//
	// +optional
	pkg string,

	// Enable data race detection.
	//
	// +optional
	race bool,

	// Arguments to pass on each go tool link invocation.
	//
	// +optional
	ldflags []string,

	// A list of additional build tags to consider satisfied during the build.
	//
	// +optional
	tags []string,

	// Remove all file system paths from the resulting executable.
	//
	// +optional
	trimpath bool,

	// Additional args to pass to the build command.
	//
	// +optional
	rawArgs []string,
) *dagger.File {
	const binaryPath = "/work/out/binary"

	args := []string{"go", "build", "-o", binaryPath}

	if race {
		args = append(args, "-race")
	}

	if len(ldflags) > 0 {
		args = append(args, "-ldflags", strings.Join(ldflags, " "))
	}

	if len(tags) > 0 {
		args = append(args, "-tags", strings.Join(tags, ","))
	}

	if trimpath {
		args = append(args, "-trimpath")
	}

	if len(rawArgs) > 0 {
		args = append(args, rawArgs...)
	}

	if pkg != "" {
		args = append(args, pkg)
	}

	return m.Container.
		WithExec(args).
		File(binaryPath)
}

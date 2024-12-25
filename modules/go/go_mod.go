// Go programming language module.
package main

import "dagger/go/internal/dagger"

// Generate Run "go generate" command.
//
// Consult "go help generate" for more information.
func (m *Go) Mod(
	// Directory with sources
	source *dagger.Directory,
) *Go {
	args := []string{"go", "mod", "download"}

	m.Container = m.Container.
		WithDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(args)

	return m
}

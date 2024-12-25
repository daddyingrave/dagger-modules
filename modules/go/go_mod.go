// Go programming language module.
package main

import "dagger/go/internal/dagger"

func (m *Go) Mod(
	// Directory with sources
	source *dagger.Directory,
) *Go {
	args := []string{"go", "mod", "download"}

	m.Container = m.Container.WithExec(args)

	return m
}

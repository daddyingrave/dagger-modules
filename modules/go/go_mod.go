// Go programming language module.
package main

func (m *Go) Mod() *Go {
	args := []string{"go", "mod", "download"}

	m.Container = m.Container.WithExec(args)

	return m
}

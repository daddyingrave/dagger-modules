// Go programming language module.
package main

// Generate Run "go generate" command.
func (m *Go) Test() *Go {
	m.Container = m.Container.
		WithExec([]string{"go", "test", "-coverprofile", "cover.txt", "./..."}).
		WithExec([]string{"mkdir", "reports"}).
		WithExec([]string{"mv", "cover.txt", "reports/cover.txt"})

	return m
}

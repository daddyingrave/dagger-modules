// Go programming language module.
package main

// Generate Run "go generate" command.
//
// Consult "go help generate" for more information.
func (m *Go) Generate(
	// Packages (or files) to run "go generate" on.
	//
	// +optional
	packages []string,

	// A regular expression to select directives whose full original source text (excluding any trailing spaces and final newline) matches the expression.
	//
	// +optional
	run string,

	// A regular expression to suppress directives whose full original source text (excluding any trailing spaces and final newline) matches the expression.
	//
	// +optional
	skip string,

	// TODO: add -v, -n and -x flags
) *Go {
	args := []string{"go", "generate"}

	if run != "" {
		args = append(args, "-run", run)
	}

	if skip != "" {
		args = append(args, "-skip", skip)
	}

	if len(packages) > 0 {
		args = append(args, packages...)
	}

	m.Container = m.Container.
		WithExec(args)

	return m
}

// A generated module for Test functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"fmt"

	"dagger/test/internal/dagger"
)

type Test struct{}

func (m *Test) TestBuild(ctx context.Context, buildDir *dagger.Directory) error {
	c, err := dag.
		GoBuild(dagger.GoBuildOpts{}).
		Build(buildDir).
		File("/src/build/test").
		Sync(ctx)
	if err != nil {
		return err
	}

	out, err := dag.Container().From("alpine").WithFile("/app", c).WithExec([]string{"/app"}).Stderr(ctx)
	if err != nil {
		return err
	}

	if out != "hello world123" {
		return fmt.Errorf("unexpected output: wanted \"hello world\", got %q", out)
	}

	return nil
}

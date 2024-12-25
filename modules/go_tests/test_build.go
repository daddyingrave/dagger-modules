package main

import (
	"context"
	"fmt"

	"dagger/go/tests/internal/dagger"
)

func (m *GoTests) TestBuild(ctx context.Context, source *dagger.Directory) error {
	binary, err := dag.Go().
		Build(source).
		Sync(ctx)
	if err != nil {
		return err
	}

	out, err := dag.
		Container().
		From("alpine").
		WithFile("/app", binary).
		WithExec([]string{"/app"}).
		Stderr(ctx)
	if err != nil {
		return err
	}

	if out != "hello" {
		return fmt.Errorf("unexpected output: wanted \"hello\", got %q", out)
	}

	return nil
}

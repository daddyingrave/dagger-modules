package main

import (
	"context"
	"fmt"
	"strings"

	"dagger/go/tests/internal/dagger"
)

func (m *GoTests) TestGenerate(ctx context.Context, source *dagger.Directory) error {
	content, err := dag.
		Go().
		Generate(source, dagger.GoGenerateOpts{Packages: []string{"./..."}}).
		Container().
		File("world").
		Contents(ctx)
	if err != nil {
		return err
	}

	if !strings.Contains(content, "hello") {
		return fmt.Errorf("unexpected output to contain \"hello\", got %q", content)
	}

	return nil
}

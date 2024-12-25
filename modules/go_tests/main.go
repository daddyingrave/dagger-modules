package main

import (
	"context"

	"dagger/go/tests/internal/dagger"
	"example.com/hello/pkg/hello"

	"github.com/sourcegraph/conc/pool"
)

type GoTests struct{}

// All executes all tests.
func (m *GoTests) All(ctx context.Context, source *dagger.Directory) error {
	hello.Hell()

	p := pool.New().WithErrors().WithContext(ctx)
	p.Go(func(ctx context.Context) error { return m.TestBuild(ctx, source) })
	p.Go(func(ctx context.Context) error { return m.TestGenerate(ctx, source) })

	return p.Wait()
}

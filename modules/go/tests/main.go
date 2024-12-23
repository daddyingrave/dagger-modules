package main

import (
	"context"
	"fmt"

	"dagger/go/tests/internal/dagger"
)

type Tests struct{}

// All executes all tests.
func (m *Tests) All(ctx context.Context, source *dagger.Directory) error {
	err := m.Build(ctx, source)
	if err != nil {
		return err
	}
	//p.Go(m.ExecBuild)
	//p.Go(m.ExecTest)
	//p.Go(m.Source)
	return nil
}

func (m *Tests) Build(ctx context.Context, source *dagger.Directory) error {
	const platform = "darwin/arm64/v7"

	binary, err := dag.Go().
		Build(source).
		Sync(ctx)
	if err != nil {
		return err
	}

	out, err := dag.Container().From("alpine").WithFile("/app", binary).WithExec([]string{"/app"}).Stderr(ctx)
	if err != nil {
		return err
	}

	if out != "hello" {
		return fmt.Errorf("unexpected output: wanted \"hello\", got %q", out)
	}

	return nil

	//p.Go(func(ctx context.Context) error {
	//	binary, err := dag.Go().
	//		WithSource(dag.CurrentModule().Source().Directory("./testdata")).
	//		Build().
	//		Sync(ctx)
	//	if err != nil {
	//		return err
	//	}
	//
	//	out, err := dag.Container().From("alpine").WithFile("/app", binary).WithExec([]string{"/app"}).Stderr(ctx)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if out != "hello" {
	//		return fmt.Errorf("unexpected output: wanted \"hello\", got %q", out)
	//	}
	//
	//	return nil
	//})
	//
	//p.Go(func(ctx context.Context) error {
	//	binary, err := dag.Go().
	//		WithSource(dag.CurrentModule().Source().Directory("./testdata")).
	//		Build(dagger.GoWithSourceBuildOpts{
	//			Ldflags: []string{"-X", "main.version=1.0.0"},
	//		}).
	//		Sync(ctx)
	//	if err != nil {
	//		return err
	//	}
	//
	//	out, err := dag.Container().From("alpine").WithFile("/app", binary).WithExec([]string{"/app", "version"}).Stderr(ctx)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if out != "1.0.0" {
	//		return fmt.Errorf("unexpected output: wanted \"1.0.0\", got %q", out)
	//	}
	//
	//	return nil
	//})
}

//func (m *Tests) ExecBuild(ctx context.Context) error {
//	ctr, err := dag.Go().
//		WithSource(dag.CurrentModule().Source().Directory("./testdata")).
//		Exec([]string{"go", "build", "-o", "/app", "."}).
//		Sync(ctx)
//	if err != nil {
//		return err
//	}
//
//	out, err := ctr.WithExec([]string{"/app"}).Stderr(ctx)
//	if err != nil {
//		return err
//	}
//
//	if out != "hello" {
//		return fmt.Errorf("unexpected output: wanted \"hello\", got %q", out)
//	}
//
//	return nil
//}
//
//func (m *Tests) ExecTest(ctx context.Context) error {
//	ctr, err := dag.Go().
//		WithSource(dag.CurrentModule().Source().Directory("./testdata")).
//		Exec([]string{"go", "test", "-v"}).
//		Sync(ctx)
//	if err != nil {
//		return err
//	}
//
//	out, err := ctr.Stdout(ctx)
//	if err != nil {
//		return err
//	}
//
//	if !strings.Contains(out, "hello") {
//		return fmt.Errorf("unexpected output to contain \"hello\", got %q", out)
//	}
//
//	return nil
//}
//
//func (m *Tests) Source(ctx context.Context) error {
//	withSource := dag.Go().
//		WithSource(dag.CurrentModule().Source().Directory("./testdata"))
//
//	out, err := withSource.Source().Entries(ctx)
//	if err != nil {
//		return err
//	}
//
//	if !reflect.DeepEqual(out, []string{"go.mod", "main.go", "main_test.go"}) {
//		return fmt.Errorf("unexpected output, got %#v", out)
//	}
//
//	return nil
//}
//
//func (m *Tests) Generate(ctx context.Context) error {
//	out, err := dag.Go().
//		WithSource(dag.CurrentModule().Source().Directory("./testdata")).
//		Generate(dagger.GoWithSourceGenerateOpts{Packages: []string{"./..."}}).
//		Source().
//		File("world").
//		Contents(ctx)
//	if err != nil {
//		return err
//	}
//
//	if !strings.Contains(out, "hello") {
//		return fmt.Errorf("unexpected output to contain \"hello\", got %q", out)
//	}
//
//	return nil
//}

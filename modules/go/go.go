package main

import (
	"fmt"

	"dagger/go/internal/dagger"
	"github.com/containerd/platforms"
)

// defaultImageRepository is used when no image is specified.
const defaultImageRepository = "golang"

const workdir = "/work/src"

type Go struct {
	Container *dagger.Container
}

func New(
	source *dagger.Directory,

	// Version (image tag) to use from the official image repository as a base container.
	//
	// +optional
	version string,

	// Custom container to use as a base container.
	//
	// +optional
	container *dagger.Container,

	// Disable mounting cache volumes.
	//
	// +optional
	disableCache bool,
) *Go {
	if container == nil {
		if version == "" {
			version = "latest"
		}

		container = dag.Container().From(fmt.Sprintf("%s:%s", defaultImageRepository, version))
	}

	container = container.
		WithDirectory("/src", source).
		WithWorkdir("/src")

	m := &Go{
		Container: container,
	}

	if !disableCache {
		m = m.
			WithModuleCache(dag.CacheVolume("go-mod"), nil, "").
			WithBuildCache(dag.CacheVolume("go-build"), nil, "")
	}

	return m
}

// WithEnvVariable Set an environment variable.
func (m *Go) WithEnvVariable(
	// The name of the environment variable (e.g., "HOST").
	name string,

	// The value of the environment variable (e.g., "localhost").
	value string,

	// Replace `${VAR}` or $VAR in the value according to the current environment
	// variables defined in the container (e.g., "/opt/bin:$PATH").
	//
	// +optional
	expand bool,
) *Go {
	m.Container = m.Container.WithEnvVariable(
		name,
		value,
		dagger.ContainerWithEnvVariableOpts{
			Expand: expand,
		},
	)

	return m
}

// WithServiceBinding Establish a runtime dependency on a service.
func (m *Go) WithServiceBinding(
	// A name that can be used to reach the service from the container.
	alias string,

	// Identifier of the service container.
	service *dagger.Service,
) *Go {
	m.Container = m.Container.WithServiceBinding(alias, service)

	return m
}

// WithPlatform Set GOOS, GOARCH and GOARM environment variables.
func (m *Go) WithPlatform(
	// Target platform in "[os]/[platform]/[version]" format (e.g., "darwin/arm64/v7", "windows/amd64", "linux/arm64").
	platform dagger.Platform,
) *Go {
	if platform == "" {
		return m
	}

	p := platforms.MustParse(string(platform))

	m.Container = m.Container.
		WithEnvVariable("GOOS", p.OS).
		WithEnvVariable("GOARCH", p.Architecture).
		With(func(c *dagger.Container) *dagger.Container {
			if p.Variant != "" {
				return c.WithEnvVariable("GOARM", p.Variant)
			}

			return c
		})

	return m
}

// WithCgoEnabled Set CGO_ENABLED environment variable to 1.
func (m *Go) WithCgoEnabled() *Go {
	m.Container = m.Container.WithEnvVariable("CGO_ENABLED", "1")

	return m
}

// WithCgoDisabled Set CGO_ENABLED environment variable to 0.
func (m *Go) WithCgoDisabled() *Go {
	m.Container = m.Container.WithEnvVariable("CGO_ENABLED", "0")

	return m
}

// WithModuleCache Mount a cache volume for Go module cache.
func (m *Go) WithModuleCache(
	cache *dagger.CacheVolume,

	// Identifier of the directory to use as the cache volume's root.
	//
	// +optional
	source *dagger.Directory,

	// Sharing mode of the cache volume.
	//
	// +optional
	sharing dagger.CacheSharingMode,
) *Go {
	m.Container = m.Container.WithMountedCache(
		"/go/pkg/mod",
		cache,
		dagger.ContainerWithMountedCacheOpts{
			Source:  source,
			Sharing: sharing,
		},
	)

	return m
}

// WithBuildCache Mount a cache volume for Go build cache.
func (m *Go) WithBuildCache(
	cache *dagger.CacheVolume,

	// Identifier of the directory to use as the cache volume's root.
	//
	// +optional
	source *dagger.Directory,

	// Sharing mode of the cache volume.
	//
	// +optional
	sharing dagger.CacheSharingMode,
) *Go {
	m.Container = m.Container.WithMountedCache(
		"/root/.cache/go-build",
		cache,
		dagger.ContainerWithMountedCacheOpts{
			Source:  source,
			Sharing: sharing,
		},
	)

	return m
}

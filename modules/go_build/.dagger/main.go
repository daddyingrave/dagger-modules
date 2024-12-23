package main

import (
	"fmt"

	"dagger/go-build/internal/dagger"
)

const (
	// defaultImageRepository is used when no image is specified.
	defaultImageRepository = "golang"
	defaultImageTag        = "1.23.4-alpine3.21"
	workdir                = "/work/src"
	goModCacheKey          = "go-mod"
	goBuildCacheKey        = "go-build"
)

type GoBuild struct {
	Container *dagger.Container
}

func New(
	// Version (image tag) to use from the official image repository as a base container.
	//
	// +optional
	// +default="1.22.10-alpine3.21"
	version string,

	// Version (image tag) to use from the official image repository as a base container.
	//
	// +optional
	// +default="golang"
	repository string,

	// Custom container to use as a base container.
	//
	// +optional
	container *dagger.Container,

	// Disable mounting cache volumes.
	//
	// +optional
	// +default="false"
	disableCache bool,
) *GoBuild {
	if container == nil {
		//if version == "" {
		version = defaultImageTag
		//}
		if repository == "" {
			repository = defaultImageRepository
		}

		container = dag.Container().From(fmt.Sprintf("%s:%s", repository, version))
	}

	m := &GoBuild{
		Container: container,
	}

	if !disableCache {
		m = m.
			WithModuleCache(dag.CacheVolume(goModCacheKey), nil, "").
			WithBuildCache(dag.CacheVolume(goBuildCacheKey), nil, "")
	}

	return m
}

// Mount a cache volume for Go module cache.
func (m *GoBuild) WithModuleCache(
	cache *dagger.CacheVolume,

	// Identifier of the directory to use as the cache volume's root.
	//
	// +optional
	source *dagger.Directory,

	// Sharing mode of the cache volume.
	//
	// +optional
	sharing dagger.CacheSharingMode,
) *GoBuild {
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

// Mount a cache volume for Go build cache.
func (m *GoBuild) WithBuildCache(
	cache *dagger.CacheVolume,

	// Identifier of the directory to use as the cache volume's root.
	//
	// +optional
	source *dagger.Directory,

	// Sharing mode of the cache volume.
	//
	// +optional
	sharing dagger.CacheSharingMode,
) *GoBuild {
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

// WithCGOEnabled Sets CGO_ENABLED env variable to 1.
// Disabled by default
func (m *GoBuild) WithCGOEnabled() *GoBuild {
	m.Container = m.Container.WithEnvVariable("CGO_ENABLED", "1")
	return m
}

func (m *GoBuild) Build(source *dagger.Directory) *dagger.Container {
	return m.Container.
		WithExec([]string{"apk", "--no-cache", "add", "build-base", "bash"}).
		WithEnvVariable("GOPATH", "/go").
		WithEnvVariable("CGO_CFLAGS", "-D_LARGEFILE64_SOURCE").
		WithDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"go", "mod", "download"}).
		WithExec([]string{"go", "build", "-o", "build/"})
}

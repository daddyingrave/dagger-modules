package main

import (
	"fmt"
	"strings"

	"example.com/hello/pkg/hello"
)

type HelloTests struct{}

func (m *HelloTests) DevilTest() error {
	numOfDevilsInMe := containsDevilCount(hello.Hell())
	if numOfDevilsInMe < 2 {
		return fmt.Errorf("i feel obsessed! You're lying me")
	}

	return nil
}

func containsDevilCount(maybeWithADevil string) int {
	return strings.Count(maybeWithADevil, "👹")
}

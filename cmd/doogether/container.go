package main

import (
	"go/internal/pkg/api/container"

	"go.uber.org/dig"
)

func InitContainer() *dig.Container {
	c := dig.New()

	if err := container.BuildConfigContainer(c); err != nil {
		panic(err)
	}

	if err := container.BuildRepositoryContainer(c); err != nil {
		panic(err)
	}

	if err := container.BuildUsecaseContainer(c); err != nil {
		panic(err)
	}

	if err := container.BuildHanlerContainer(c); err != nil {
		panic(err)
	}

	return c
}

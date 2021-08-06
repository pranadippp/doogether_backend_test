package container

import (
	"go/internal/pkg/service"

	"go.uber.org/dig"
)

func BuildRepositoryContainer(container *dig.Container) error {
	if err := container.Provide(service.NewUserServices); err != nil {
		return err
	}

	if err := container.Provide(service.NewSessionServices); err != nil {
		return err
	}

	return nil
}

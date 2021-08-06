package container

import (
	"go/internal/pkg/api/app/context/mst_user"
	"go/internal/pkg/api/app/context/session"

	"go.uber.org/dig"
)

func BuildUsecaseContainer(container *dig.Container) error {
	if err := container.Provide(mst_user.NewUserUseCase); err != nil {
		return err
	}

	if err := container.Provide(session.NewSessionUseCase); err != nil {
		return err
	}

	return nil
}

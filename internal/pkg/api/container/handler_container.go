package container

import (
	"go/internal/pkg/api/app/handler/mst_user"
	"go/internal/pkg/api/app/handler/session"

	"go.uber.org/dig"
)

func BuildHanlerContainer(container *dig.Container) error {
	if err := container.Provide(mst_user.NewUserHandler); err != nil {
		return err
	}

	if err := container.Provide(session.NewSessionHandler); err != nil {
		return err
	}

	return nil
}

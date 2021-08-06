package container

import (
	"go/internal/pkg/config"

	"github.com/gin-gonic/gin"

	"go.uber.org/dig"
)

func BuildConfigContainer(container *dig.Container) error {

	//debug mode
	if err := container.Provide(config.NewAppConfig); err != nil {
		return err
	}

	if err := container.Provide(config.NewDatabase); err != nil {
		return err
	}

	if err := container.Provide(func() *gin.Engine {
		return gin.Default()
	}); err != nil {
		return err
	}

	return nil
}

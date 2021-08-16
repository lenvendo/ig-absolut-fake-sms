package container

import (
	"gorm.io/gorm"
	"go.uber.org/zap"

	"github.com/lenvendo/ig-absolut-fake-sms/lib/config"
	"github.com/lenvendo/ig-absolut-fake-sms/service"
	"github.com/lenvendo/ig-absolut-fake-sms/web"
)

type Container struct {
	Config config.Config

	Gorm   *gorm.DB
	Logger *zap.Logger

	WorkerService *service.WorkerService
	CodesHandler  *web.CodesHandler
}

func NewContainer() *Container {
	c := &Container{}
	return c.Init()
}

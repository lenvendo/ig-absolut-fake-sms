package container

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	gormProm "gorm.io/plugin/prometheus"
	"go.uber.org/zap"

	"github.com/lenvendo/ig-absolut-fake-sms/service"
	"github.com/lenvendo/ig-absolut-fake-sms/lib/config"
	"github.com/lenvendo/ig-absolut-fake-sms/lib/db"
	"github.com/lenvendo/ig-absolut-fake-sms/lib/log"
	learn "github.com/lenvendo/ig-absolut-fake-sms/web"
)

func (c *Container) Init() *Container {
	println("Container building ...")

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	c.Config = config.InitConfig()
	cfg := c.Config
	logger := c.initLogger(cfg)
	c.initGorm(cfg)

	c.WorkerService = service.NewWorkerService(c.Gorm, c.Logger)
	c.CodesHandler = learn.NewCodesHandler(c.Gorm, c.Logger)

	logger.Info("Container is ready")

	return c
}

func (c *Container) initLogger(cfg config.Config) *zap.Logger {
	logger, err := log.NewLoggerFromConfig(cfg.Logger, cfg.IsDebug)
	if err != nil {
		fmt.Printf("unable to initialize logger: %v\n", err)
		os.Exit(1)
	}

	c.Logger = logger
	return c.Logger
}
func (c *Container) initGorm(cfg config.Config) {
	gormDB, err := db.NewGorm(*cfg.Database, cfg.IsDebug)
	if err != nil {
		c.Logger.Fatal("gorm init failed", zap.Error(err))
	}

	err = gormDB.Use(gormProm.New(gormProm.Config{
		DBName:          cfg.Database.DatabaseName,
		RefreshInterval: 5,
	}))
	if err != nil {
		c.Logger.Fatal("gorm init failed", zap.Error(err))
	}

	c.Gorm = gormDB
}

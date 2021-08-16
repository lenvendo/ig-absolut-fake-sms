package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGorm(cfg PostgreSQLConfig, isDebug bool) (*gorm.DB, error) {
	_gorm, err := gorm.Open(postgres.New(postgres.Config{
		DSN: cfg.GetDsn(),
	}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		},
	)
	if err != nil {
		return nil, err
	}

	sqlDB, err := _gorm.DB()
	if err != nil {
		return nil, err
	}

	// connection pool settings
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if isDebug {
		return _gorm.Debug(), nil
	}
	return _gorm, nil
}

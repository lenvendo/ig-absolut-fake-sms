package db

import (
	"fmt"
	"time"
)

type PostgreSQLConfig struct {
	Server          string        `envconfig:"server"`
	Port            string        `envconfig:"port"`
	User            string        `envconfig:"user"`
	Password        string        `envconfig:"password"`
	DatabaseName    string        `envconfig:"db_name"`
	DialTimeout     string        `envconfig:"dial_timeout" default:"2s"`
	IOTimeout       string        `envconfig:"io_timeout" default:"5s"`
	MaxOpenConns    int           `envconfig:"maxopenconns"`
	MaxIdleConns    int           `envconfig:"maxidleconns" default:"2"`
	Driver          string        `envconfig:"driver" default:"mysql"`
	ConnMaxLifetime time.Duration `envconfig:"connmaxlifetime" default:"10m"`
	MigrationsTable string        `envconfig:"migrations_table" default:"migrations"`
}

func (cfg *PostgreSQLConfig) GetDsn() string {

	return fmt.Sprintf("port=%s host=%s user=%s password=%s database=%s",
		cfg.Port,
		cfg.Server,
		cfg.User,
		cfg.Password,
		cfg.DatabaseName,
	)
}

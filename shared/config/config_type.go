package config

import (
	"time"
)

type DefaultConfig struct {
	Apps            Apps            `mapstructure:"apps"`
	Server          Server          `mapstructure:"server"`
	Database        Database        `mapstructure:"database"`
	DatabaseSeeder  DatabaseSeeder  `mapstructure:"databaseSeeder"`
	PasswordHashing PasswordHashing `mapstructure:"passwordHashing"`
}

type Apps struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Datasource struct {
	Url               string        `mapstructure:"url"`
	Port              string        `mapstructure:"port"`
	DatabaseName      string        `mapstructure:"databaseName"`
	Username          string        `mapstructure:"username"`
	Password          string        `mapstructure:"password"`
	Schema            string        `mapstructure:"schema"`
	ConnectionTimeout time.Duration `mapstructure:"connectionTimeout"`
	MaxIdleConnection int           `mapstructure:"maxIdleConnection"`
	MaxOpenConnection int           `mapstructure:"maxOpenConnection"`
	DebugMode         bool          `mapstructure:"debugMode"`
}

type SeederConfig struct {
	IsRebuildData bool `mapstructure:"isRebuildData"`
}

type Database struct {
	Postgres Datasource `mapstructure:"postgres"`
}

type DatabaseSeeder struct {
	User SeederConfig `mapstructure:"user"`
}

type PasswordHashing struct {
	HashSalt int `mapstructure:"hashSalt"`
}

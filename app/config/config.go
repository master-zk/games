package config

type Config struct {
	App       AppConfig
	Databases DatabasesConfig `mapstructure:"database"`
	Log       LogConfig
}

type DatabasesConfig struct {
	Jenny DatabaseConfig `mapstructure:"jenny"`
	Xh    DatabaseConfig `mapstructure:"xh"`
}

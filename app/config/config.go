package config

type Config struct {
	App       AppConfig
	Server    ServerConfig
	Databases DatabasesConfig `mapstructure:"database"`
	Redis     RedisConfig     `mapstructure:"redis"`
	Log       LogConfig
}

type DatabasesConfig struct {
	Jenny DatabaseConfig `mapstructure:"jenny"`
	Xh    DatabaseConfig `mapstructure:"xh"`
}

type ServerConfig struct {
	Port         string `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"readTimeout"`
	WriteTimeout int    `mapstructure:"writeTimeout"`
}

type RedisConfig struct {
	Host     string
	Port     string
	Auth     string
	Database int
}

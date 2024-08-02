package config

type Config struct {
	App    AppConfig
	Server ServerConfig
	Mysql  map[string]DatabaseConfig `mapstructure:"mysql"`
	Redis  map[string]RedisConfig    `mapstructure:"redis"`
	Log    LogConfig
}

type AppConfig struct {
	Name     string
	Key      string
	Timezone string
	Debug    bool
	Mode     string
}

type LogConfig struct {
	Channel  string
	Level    string
	Path     string
	LogStore string
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

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Pool     PoolConfig
}

type PoolConfig struct {
	MaxIdleConn int
	MaxOpenConn int
	MaxLifetime int64
}

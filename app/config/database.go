package config

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

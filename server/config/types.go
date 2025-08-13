package config

type Config struct {
	Server      ServerConfig      `koanf:"server"`
	Database    DatabaseConfig    `koanf:"database"`
	Auth        AuthConfig        `koanf:"auth"`
	ExternalAPI ExternalAPIConfig `koanf:"external_apis"`
	Logging     LoggingConfig     `koanf:"logging"`
	Cache       CacheConfig       `koanf:"cache"`
}

type ServerConfig struct {
	Host string `koanf:"host"`
	Port int    `koanf:"port"`
}

type DatabaseConfig struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	User     string `koanf:"port"`
	Password string `koanf:"password"`
	DBName   string `koanf:"dbname"`
	SSLMode  string `koanf:"sslmode"`
}

type AuthConfig struct {
	JwtSecret          string `koanf:"jwtSecret"`
	TokenExpiryMinutes int    `koanf:"tokenExpiryMinutes"`
}
type ExternalAPIConfig struct {
}
type LoggingConfig struct {
	Level string `koanf:"level"`
}

type CacheConfig struct {
	Redis RedisConfig `koanf:"redis"`
}

type RedisConfig struct {
	Enabled  bool   `koanf:"enabled"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Password string `koanf:"password"`
	DB       int    `koanf:"db"`
}

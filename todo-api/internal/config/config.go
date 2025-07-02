package config

// Config はアプリケーション全体の設定を保持します。
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Nats     NatsConfig     `mapstructure:"nats"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

type NatsConfig struct {
	URL        string   `mapstructure:"url"`
	StreamName string   `mapstructure:"streamName"`
	Subjects   []string `mapstructure:"subjects"`
}

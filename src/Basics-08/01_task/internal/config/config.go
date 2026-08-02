package config

type Config struct {
	Name        string
	Host        string            `env:"app_host"`
	Port        int               `env:"app_port"`
	Environment map[string]string `env:"environment"`
	Volumes     []string          `env:"volumes"`
}

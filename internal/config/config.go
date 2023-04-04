package config

import "time"

type Application struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"DEBUG"`
	Server   Server
}

type Server struct {
	Listen       string        `env:"API_LISTEN" envDefault:":8080"`
	ReadTimeout  time.Duration `env:"API_READ_TIMEOUT" envDefault:"30s"`
	WriteTimeout time.Duration `env:"API_WRITE_TIMEOUT" envDefault:"30s"`
}

package config

import "time"

type Application struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"DEBUG"`
	Server   Server
	Auth     Auth
}

type Server struct {
	Listen       string        `env:"API_LISTEN" envDefault:":8080"`
	ReadTimeout  time.Duration `env:"API_READ_TIMEOUT" envDefault:"30s"`
	WriteTimeout time.Duration `env:"API_WRITE_TIMEOUT" envDefault:"30s"`
}

type Auth struct {
	AccessTokenTime  time.Duration `env:"AUTH_ACCESS_TOKEN_TIME" envDefault:"30m"`
	RefreshTokenTime time.Duration `env:"AUTH_REFRESH_TOKEN_TIME" envDefault:"720h"` // default: 30 days
	Issuer           string        `env:"AUTH_ISSUER" envDefault:"auth.tikkichest.ru"`
	PrivateKeyFile   string        `env:"AUTH_PRIVATE_KEY" envDefault:"id_rsa"`
	PublicKeyFile    string        `env:"AUTH_PUBLIC_KEY" envDefault:"id_rsa.pub"`
}

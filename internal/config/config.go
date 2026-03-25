package config

import (
	"log"
	"net"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type HTTPServer struct {
	Host string `env:"HOST"`
	Port string `env:"PORT" env-required:"true"`
}

type Config struct {
	ENV    string `env:"ENV" env-default:"Prod"`
	DB_URL string `env:"DB_URL" env-required:"true"`

	JWT_SECRET string `env:"JWT_SECRET" env-required:"true"`

	HTTPServer
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	LoadEnv()

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}

func (s HTTPServer) Addr() string {
	return net.JoinHostPort(s.Host, s.Port)
}

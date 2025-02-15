package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type HTTPServer struct {
	Hostname string `yaml:"Hostname"`
	Addr     string `yaml:"Addr"`
}

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"sslmode"`
}

type Config struct {
	Env        string         `yaml:"env" env:"ENV" env-required:"true"`
	HTTPServer HTTPServer     `yaml:"http_server"`
	Database   DatabaseConfig `yaml:"database"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {

		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		if *flags == "" {
			log.Fatal("No config file specified via flag or environment variable")
		}
		configPath = *flags
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}

	if cfg.Env == "" {
		log.Fatal("Missing required ENV variable")
	}

	return &cfg
}

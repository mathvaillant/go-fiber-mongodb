package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment string `env:"ENVIRONMENT,required"`
	Version     int    `env:"VERSION,required"`
	MongoURI    string `env:"MONGO_URI,required"`
	MongoDBName string `env:"MONGO_DB_NAME,required"`
	Port        string `env:"PORT,required"`
}

func NewConfig() (*Config, error) {
	// Loading the environment variables from '.env' file.
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Unable to load .env file: %e", err)
	}

	config := &Config{} // pointer to the new instance of `Config`

	err = env.Parse(config) // Parse environment variables into `Config`

	if err != nil {
		log.Fatalf("Unable to parse environment variables: %e", err)
	}

	return config, nil
}

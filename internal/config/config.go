package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type RedisConfig struct {
	Host string
	Port string
}

func LoadConfig() (*Config, error) {
	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	viper.SetConfigFile("configs/config.yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("POSTGRES_DB", os.Getenv("POSTGRES_DB"))
	viper.SetDefault("POSTGRES_USER", os.Getenv("POSTGRES_USER"))
	viper.SetDefault("POSTGRES_PASSWORD", os.Getenv("POSTGRES_PASSWORD"))
	viper.SetDefault("POSTGRES_HOST", os.Getenv("POSTGRES_HOST"))
	viper.SetDefault("POSTGRES_PORT", os.Getenv("POSTGRES_PORT"))
	viper.SetDefault("REDIS_HOST", os.Getenv("REDIS_HOST"))
	viper.SetDefault("REDIS_PORT", os.Getenv("REDIS_PORT"))
	viper.SetDefault("SERVER_PORT", os.Getenv("SERVER_PORT"))
}

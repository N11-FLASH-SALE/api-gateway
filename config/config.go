package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	USER_SERVICE string
	SALE_SERVICE string
	API_ROUTER   string

	ACCES_KEY   string
	REFRESH_KEY string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found?")
	}

	config := Config{}
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", ":50051"))
	config.SALE_SERVICE = cast.ToString(Coalesce("SALE_SERVICE", ":50052"))
	config.API_ROUTER = cast.ToString(Coalesce("API_ROUTER", ":8080"))
	config.ACCES_KEY = cast.ToString(Coalesce("ACCES_KEY", "flashsalee"))
	config.REFRESH_KEY = cast.ToString(Coalesce("REFRESH_KEY", "OzNur"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

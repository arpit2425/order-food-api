package config

import "os"

type Config struct {
	ServerPort string
	ApiKey     string
}

func Load() Config {
	port := getenv("port", ":8080")
	apiKey := getenv("API_KEY", "apitest")
	return Config{
		ServerPort: port,
		ApiKey:     apiKey,
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

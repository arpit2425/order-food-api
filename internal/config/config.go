package config

import "os"

type Config struct {
	ServerPort string
}

func Load() Config {
	port := getenv("port", ":8080")
	return Config{
		ServerPort: port,
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

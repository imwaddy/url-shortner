package config

import "os"

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBName string
	Port   string
}

func Load() *Config {
	return &Config{
		DBUser: getEnv("DB_USER", "root"),
		DBPass: getEnv("DB_PASS", "root"),
		DBHost: getEnv("DB_HOST", "mysql:3306"),
		DBName: getEnv("DB_NAME", "shortner"),
		Port:   getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

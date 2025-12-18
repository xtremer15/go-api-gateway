package config

import (
	"os"
)

type Config struct {
	//Port on which the application will run
	Port string
	//Env that the application is running in (development, staging, production)
	Env string
	//Will hold the database connection string
	//Will have an adapter to different databases in future
	DbConnectionString string
	LogLevel           string
	//The key will be the feature name and the value will indicate if it's enabled or not
	// e.g. "NewDashboard": true
	FeatureFlags   map[string]bool
	Timeout        int
	RequestLimiter int
	BaseURL        string
}

func Load() *Config {
	return &Config{
		Port:               getEnv("PORT", "8080"),
		Env:                getEnv("DEV", "DEV"),
		DbConnectionString: getEnv("DBCREDS", "user:password@/dbname"),
		LogLevel:           getEnv("LOG_LEVEL", "INFO"),
		FeatureFlags:       make(map[string]bool),
		BaseURL:            getEnv("URL", "https://dummyjson.com"),
	}
}

func getEnv(key, defaultValue string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func setEnv(key, value string) {
	os.Setenv(key, value)
}

package config

import "os"

type Config struct {
	Port          string
	CassandraHost string
	Keyspace      string
}

// In Go, config is specified in environment variables to be configured at the deployment phase
func Load() Config {
	return Config{
		Port:          getEnv("PORT", ""),
		CassandraHost: getEnv("CASSANDRA_HOST", ""),
		Keyspace:      getEnv("CASSANDRA_KEYSPACE", ""),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

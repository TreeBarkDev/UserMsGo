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
		Port:          getEnv("PORT", "8082"),
		CassandraHost: getEnv("CASSANDRA_HOST", "127.0.0.1"),
		Keyspace:      getEnv("CASSANDRA_KEYSPACE", "demo"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

package config

import "os"

// Config structure to store env vars
type Config struct {
	GRPCport string
	HTTPport string
	RedisURL string
}

// New function to get env vars
func New() *Config {
	return &Config{
		GRPCport: getEnv("GRPC_PORT", "8088"),
		HTTPport: getEnv("HTTP_PORT", "8443"),
		RedisURL: getEnv("REDIS_URL", "0.0.0.0:6379"),
	}
}

func getEnv(envVar string, defaultVal string) string {
	value, err := os.LookupEnv(envVar)
	if err != true {
		value = defaultVal
	}
	return value
}

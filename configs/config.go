package configs

import "os"

type Config struct {
	DBUserName string
	DBPassword string
	DBHost     string
	DBName     string
	DBPort     string
	ServerHost string
	ServerPort string
	Env        string
}

var cnf *Config

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func loadConfig() *Config {
	cnf = &Config{}
	cnf.DBHost = getEnv("DB_HOST", "db")
	cnf.DBUserName = getEnv("DB_USERNAME", "spice")
	cnf.DBPassword = getEnv("DB_PASSWORD", "password")
	cnf.DBName = getEnv("DB_NAME", "default")
	cnf.DBPort = getEnv("DB_PORT", "3306")
	cnf.ServerHost = getEnv("SERVER_HOST", "localhost")
	cnf.ServerPort = getEnv("SERVER_PORT", "8080")
	cnf.Env = getEnv("ENV", "develop")
	return cnf
}

func Get() Config {
	if cnf == nil {
		cnf = loadConfig()
	}
	return *cnf
}

package config

import "os"

type Config struct {
	MYSQL_HOST     string
	MYSQL_USERNAME string
	MYSQL_PASSWORD string
	MYSQL_DB_NAME  string
	MYSQL_PORT     string
}

func GetConfig(config *Config) {
	config.MYSQL_HOST = os.Getenv("MYSQL_HOST")
	config.MYSQL_USERNAME = os.Getenv("MYSQL_USER")
	config.MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	config.MYSQL_DB_NAME = os.Getenv("MYSQL_DB_NAME")
	config.MYSQL_PORT = os.Getenv("MYSQL_PORT")
}

package options

import (
	"aurelius/libs/aurelius-mysql/db_context"
	"os"
)

func NewDbOptions() *db_context.DbOptions {
	return &db_context.DbOptions{
		Host:     defaultIfEmpty(os.Getenv("DB_HOST"), "localhost"),
		Port:     defaultIfEmpty(os.Getenv("DB_PORT"), "3308"),
		User:     defaultIfEmpty(os.Getenv("DB_USER"), "aurelius"),
		Password: defaultIfEmpty(os.Getenv("DB_PASSWORD"), "aurelius"),
		Database: defaultIfEmpty(os.Getenv("DB_DATABASE"), "aurelius"),
	}
}

func defaultIfEmpty(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

package utils

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() error {
	return godotenv.Load(".env")
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
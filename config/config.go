package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func ReloadEnv() error {
	err := godotenv.Load("config/dev.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return err
	}
	return nil
}

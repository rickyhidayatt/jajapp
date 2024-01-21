package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ReloadEnv() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return err
	}
	return nil
}

func getFromFile(key string) string {
	var configEnv map[string]string
	content, err := os.ReadFile("./config/config.json")
	if err != nil {
		log.Fatal(err)
		return ""
	}

	err = json.Unmarshal(content, &configEnv)
	if err != nil {
		return ""
	}

	value, ok := configEnv[key]
	if !ok {
		return ""
	}

	return value
}

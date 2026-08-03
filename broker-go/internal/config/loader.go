package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {

	paths := []string{
		"../../.env",
		"../.env",
		".env",
	}

	for _, p := range paths {

		if _, err := os.Stat(p); err == nil {

			if err := godotenv.Load(p); err != nil {
				return err
			}

			fmt.Println("Loaded .env:", p)

			return nil
		}
	}

	return fmt.Errorf(".env not found")
}

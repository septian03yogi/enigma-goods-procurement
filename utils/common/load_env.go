package common

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() (err error) {
	err = godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file")
	}
	return
}

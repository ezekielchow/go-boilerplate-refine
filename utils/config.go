package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	APP_PORT           string `json:"APP_PORT"`
	POSTGRES_HOST      string `json:"POSTGRES_HOST"`
	POSTGRES_USER      string `json:"POSTGRES_USER"`
	POSTGRES_PASSWORD  string `json:"POSTGRES_PASSWORD"`
	POSTGRES_DB        string `json:"POSTGRES_DB"`
	POSTGRES_PORT      string `json:"POSTGRES_PORT"`
	TOKEN_EXPIRE_HOURS string `json:"TOKEN_EXPIRE_HOURS"`
	DISABLE_AUTH       string `json:"DISABLE_AUTH"`
	APP_ENV            string `json:"APP_ENV"`
	FRONTEND_API       string `json:"FRONTEND_API"`
	FB_APP_ID          string `json:"FB_APP_ID"`
	FB_APP_SECRET      string `json:"FB_APP_SECRET"`
	DSN                string `json:"DSN"`
}

func LoadEnv() (config *Config, err error) {

	envPairs := map[string]interface{}{}

	for _, env := range os.Environ() {

		s := strings.SplitN(env, "=", 2)

		if len(s) == 2 {
			envPairs[s[0]] = s[1]
		} else {
			fmt.Printf("\nEnv %s is missing", env)
		}
	}

	b, err := json.Marshal(envPairs)

	if err != nil {
		return
	}

	err = json.Unmarshal(b, &config)

	return
}

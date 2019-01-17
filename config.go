package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	AccessToken       string `envconfig:"OAUTH_ACCESS_TOKEN" required:"true"`
	UserToken         string `envconfig:"USER_TOKEN" required:"true"`
	BotAccessToken    string `envconfig:"BOT_ACCESS_TOKEN" required:"true"`
	VerificationToken string `envconfig:"VERIFICATION_TOKEN" required:"true"`
	BotID             string `envconfig:"BOT_ID" required:"true"`
	ChannelID         string `envconfig:"CHANNEL_ID" required:"true"`
	AdminChannelID    string `envconfig:"ADMIN_CHANNEL_ID" required:"false"`
}

func loadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error: Failed to load env file")
	}
	loadEnvconfig()
}

func loadEnvconfig() {
	if err := envconfig.Process("", &Env); err != nil {
		log.Printf("error: Failed to process env var: %s", err)
	}
}

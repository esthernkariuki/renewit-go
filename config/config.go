package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func DarajaConsumerKey() string {
	return os.Getenv("DARAJA_CONSUMER_KEY")
}

func DarajaConsumerSecret() string {
	return os.Getenv("DARAJA_CONSUMER_SECRET")
}

func DarajaShortcode() string {
	return os.Getenv("DARAJA_SHORTCODE")
}

func DarajaPasskey() string {
	return os.Getenv("DARAJA_PASSKEY")
}

func DarajaCallbackURL() string {
	return os.Getenv("DARAJA_CALLBACK_URL")
}

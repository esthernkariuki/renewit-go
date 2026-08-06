package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using environment variables")
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

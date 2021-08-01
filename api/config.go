package api

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBConnectionString string `default:"file:grack.db?_fk=1" split_words:"true"`
	BindAddr           string `default:":8081" split_words:"true"`

	JwtSigningSecret string `default:"grack" split_words:"true"`

	GithubClientId     string `default:"" split_words:"true"`
	GithubClientSecret string `default:"" split_words:"true"`

	DiscordClientId     string `default:"" split_words:"true"`
	DiscordClientSecret string `default:"" split_words:"true"`
	DiscordCallbackUrl  string `default:"http://localhost:8081/oauth/discord/callback" split_words:"true"`
}

var config Config

func LoadConfig() error {
	godotenv.Load()

	err := envconfig.Process("grackdb", &config)
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	return nil
}

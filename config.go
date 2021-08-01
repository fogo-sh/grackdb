package grackdb

import (
	"fmt"
	"os"

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

var AppConfig Config

func LoadConfig() error {
	envfile, ok := os.LookupEnv("GRACKDB_DOTENV_PATH")
	if ok {
		godotenv.Load(envfile)
	} else {
		godotenv.Load(".env")
	}

	err := envconfig.Process("grackdb", &AppConfig)
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	return nil
}

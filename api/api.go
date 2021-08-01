package api

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fogo-sh/grackdb"
	"github.com/fogo-sh/grackdb/ent"
	"github.com/fogo-sh/grackdb/ent/migrate"
	"github.com/fogo-sh/grackdb/graphql"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func playgroundHandler() gin.HandlerFunc {
	playgroundHandlerInstance := playground.Handler("GrackDB", "/query")

	return func(c *gin.Context) {
		playgroundHandlerInstance.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlHandler(server *handler.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	}
}

var entClient *ent.Client

func StartApi() error {
	err := grackdb.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	githubOauthConfig = &oauth2.Config{
		ClientID:     grackdb.AppConfig.GithubClientId,
		ClientSecret: grackdb.AppConfig.GithubClientSecret,
		Scopes:       []string{"read:user"},
		Endpoint:     github.Endpoint,
	}
	discordOauthConfig = &oauth2.Config{
		ClientID:     grackdb.AppConfig.DiscordClientId,
		ClientSecret: grackdb.AppConfig.DiscordClientSecret,
		Scopes:       []string{"identify"},
		RedirectURL:  grackdb.AppConfig.DiscordCallbackUrl,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discordapp.com/api/oauth2/authorize",
			TokenURL: "https://discordapp.com/api/oauth2/token",
		},
	}

	entClient, err = ent.Open(dialect.SQLite, grackdb.AppConfig.DBConnectionString)
	if err != nil {
		return fmt.Errorf("error opening ent client: %w", err)
	}
	if err := entClient.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		return fmt.Errorf("error migrating database: %w", err)
	}

	app := gin.Default()

	srv := handler.NewDefaultServer(graphql.NewSchema(entClient))

	app.GET("/", playgroundHandler())
	app.POST("/query", graphqlHandler(srv))

	app.GET("/oauth/github/auth", handleGithubAuth)
	app.GET("/oauth/github/callback", handleGithubCallback)

	app.GET("/oauth/discord/auth", handleDiscordAuth)
	app.GET("/oauth/discord/callback", handleDiscordCallback)

	if err := app.Run(grackdb.AppConfig.BindAddr); err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}

	return nil
}

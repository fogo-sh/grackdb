package api

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/executor"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fogo-sh/grackdb"
	"github.com/fogo-sh/grackdb/ent"
	"github.com/fogo-sh/grackdb/graphql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func jwtAuthMiddleware(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		authHeader := c.GetHeader("Authorization")
		jwtCookie, err := c.Cookie("jwt")
		if err == nil {
			token = jwtCookie
		} else if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}

		if token == "" {
			c.Set("user", (*ent.User)(nil))
			return
		}

		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
			}

			return []byte(grackdb.AppConfig.JwtSigningSecret), nil
		})

		if err != nil {
			fmt.Printf("Error parsing JWT: %s\n", err)
			c.Set("user", (*ent.User)(nil))
			return
		}

		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			userId, ok := claims["sub"].(float64)
			if !ok {
				fmt.Printf("Got non-float64 user ID %s\n", claims["sub"])
				c.Set("user", (*ent.User)(nil))
				return
			}
			user, err := client.User.Get(c, int(userId))
			if err != nil {
				fmt.Printf("Error fetching JWT user: %s\n", err)
				c.Set("user", (*ent.User)(nil))
			} else {
				c.Set("user", user)
			}
		} else {
			fmt.Printf("Error parsing JWT: %s\n", err)
			c.Set("user", (*ent.User)(nil))
			return
		}
	}
}

func ginContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
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
	if err = entClient.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("error migrating database: %w", err)
	}

	app := gin.Default()
	app.Use(jwtAuthMiddleware(entClient))
	app.Use(ginContextToContextMiddleware())

	executor.New(graphql.NewSchema(entClient))

	srv := handler.NewDefaultServer(graphql.NewSchema(entClient))
	srv.Use(entgql.Transactioner{TxOpener: entClient})

	app.Static("/assets", "./frontend/dist/assets")

	app.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})
	// app.StaticFile("/", "./frontend/dist/index.html")

	app.GET("/playground", playgroundHandler())
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

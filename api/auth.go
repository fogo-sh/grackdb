package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"

	"github.com/fogo-sh/grackdb"
	"github.com/fogo-sh/grackdb/ent"
	"github.com/fogo-sh/grackdb/ent/discordaccount"
	"github.com/fogo-sh/grackdb/ent/githubaccount"
)

var discordOauthConfig *oauth2.Config
var githubOauthConfig *oauth2.Config

func handleGithubAuth(c *gin.Context) {
	c.Redirect(
		http.StatusTemporaryRedirect,
		githubOauthConfig.AuthCodeURL("state"),
	)
}

func handleGithubCallback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	if state != "state" {
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"status":  "error",
				"message": "invalid state",
			},
		)
		return
	}

	token, err := githubOauthConfig.Exchange(c, code)
	if err != nil {
		fmt.Printf("Error exchanging code: %s\n", err)
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"status":  "error",
				"message": "invalid code",
			},
		)
		return
	}

	ghClient := github.NewClient(githubOauthConfig.Client(c, token))

	user, _, err := ghClient.Users.Get(c, "")
	if err != nil {
		fmt.Printf("error fetching details for github user: %s\n", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{},
		)
		return
	}

	ghAccount, err := entClient.GithubAccount.Query().
		Where(
			githubaccount.UsernameEQ(*user.Login),
			githubaccount.HasOwner(),
		).
		WithOwner().
		First(c)

	if err != nil {
		_, ok := err.(*ent.NotFoundError)
		if ok {
			c.Data(
				http.StatusForbidden,
				"text/plain;charset=UTF-8",
				[]byte("You aren't part of the GrackDB - have another member add you and then try again."),
			)
			return
		}
		fmt.Printf("error fetching github account: %s\n", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{},
		)
		return
	}

	unsignedToken := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		jwt.MapClaims{
			"iat": time.Now().UTC().Unix(),
			"sub": ghAccount.Edges.Owner.ID,
		},
	)

	tokenString, err := unsignedToken.SignedString([]byte(grackdb.AppConfig.JwtSigningSecret))

	if err != nil {
		fmt.Printf("error signing token: %s\n", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{},
		)
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("jwt", tokenString, 0, "/", "", false, false)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func handleDiscordAuth(c *gin.Context) {
	c.Redirect(
		http.StatusTemporaryRedirect,
		discordOauthConfig.AuthCodeURL("state"),
	)
}

func handleDiscordCallback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	if state != "state" {
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"status":  "error",
				"message": "invalid state",
			},
		)
		return
	}

	token, err := discordOauthConfig.Exchange(c, code)
	if err != nil {
		fmt.Printf("Error exchanging code: %s\n", err)
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"status":  "error",
				"message": "invalid code",
			},
		)
		return
	}

	discordClient, err := discordgo.New(fmt.Sprintf("Bearer %s", token.AccessToken))
	if err != nil {
		fmt.Printf("error creating discord client: %s\n", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{},
		)
		return
	}

	discordUser, err := discordClient.User("@me")
	if err != nil {
		fmt.Printf("error fetching discord user: %s\n", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{},
		)
		return
	}

	discordAccountEntity, err := entClient.DiscordAccount.Query().
		Where(
			discordaccount.DiscordIDEQ(discordUser.ID),
			discordaccount.HasOwner(),
		).
		WithOwner().
		First(c)

	if err != nil {
		_, ok := err.(*ent.NotFoundError)
		if ok {
			c.Data(
				http.StatusForbidden,
				"text/plain;charset=UTF-8",
				[]byte("You aren't part of the GrackDB - have another member add you and then try again."),
			)
			return
		}
		fmt.Printf("error fetching github account: %s\n", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{},
		)
		return
	}

	unsignedToken := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		jwt.MapClaims{
			"iat": time.Now().UTC().Unix(),
			"sub": discordAccountEntity.Edges.Owner.ID,
		},
	)

	tokenString, err := unsignedToken.SignedString([]byte(grackdb.AppConfig.JwtSigningSecret))

	if err != nil {
		fmt.Printf("error signing token: %s\n", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{},
		)
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("jwt", tokenString, 0, "/", "", false, false)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

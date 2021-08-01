package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fogo-sh/grackdb/ent"
	"github.com/fogo-sh/grackdb/ent/githubaccount"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
)

// var discordOauthConfig oauth2.Config
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

	tokenString, err := unsignedToken.SignedString([]byte(config.JwtSigningSecret))

	if err != nil {
		fmt.Printf("error signing token: %s\n", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{},
		)
		return
	}

	c.Data(
		http.StatusOK,
		"text/plain;charset=UTF-8",
		[]byte(fmt.Sprintf("Welcome to GrackDB!\n\nYour auth token is:\n%s\n\nYou can use this to query restricted API fields (and in the future, to make changes to the DB!)", tokenString)),
	)

}

package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/fogo-sh/grackdb/ent"
	"github.com/fogo-sh/grackdb/ent/githuborganizationmember"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:grack.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	riley, err := client.User.Create().SetUsername("nint8835").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test user: %v", err)
	}

	jack, err := client.User.Create().SetUsername("jackharrhy").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test user: %v", err)
	}

	_, err = client.DiscordAccount.Create().
		SetDiscordID("106162668032802816").
		SetUsername("nint8835").
		SetDiscriminator("0001").
		SetOwner(riley).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test discord account: %v", err)
	}

	_, err = client.DiscordAccount.Create().
		SetDiscordID("178958252820791296").
		SetUsername("<i>jack arthur null</i>").
		SetDiscriminator("7539").
		SetOwner(jack).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test discord account: %v", err)
	}

	rileyGithub, err := client.GithubAccount.Create().
		SetUsername("nint8835").
		SetOwner(riley).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test github account: %v", err)
	}

	jackGithub, err := client.GithubAccount.Create().
		SetUsername("jackharrhy").
		SetOwner(jack).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test github account: %v", err)
	}

	fogoSh, err := client.GithubOrganization.Create().
		SetName("fogo-sh").
		SetDisplayName("fogo.sh").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test github org: %v", err)
	}

	_, err = client.GithubOrganizationMember.Create().
		SetRole(githuborganizationmember.RoleAdmin).
		SetOrganization(fogoSh).
		SetAccount(rileyGithub).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test github org membership: %v", err)
	}

	_, err = client.GithubOrganizationMember.Create().
		SetRole(githuborganizationmember.RoleAdmin).
		SetOrganization(fogoSh).
		SetAccount(jackGithub).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test github org membership: %v", err)
	}
}

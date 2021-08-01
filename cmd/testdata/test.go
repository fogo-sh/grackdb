package main

import (
	"context"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"github.com/fogo-sh/grackdb"
	"github.com/fogo-sh/grackdb/ent"
	"github.com/fogo-sh/grackdb/ent/githuborganizationmember"
	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/projectcontributor"
	_ "github.com/fogo-sh/grackdb/ent/runtime"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := grackdb.LoadConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}
	client, err := ent.Open(dialect.SQLite, grackdb.AppConfig.DBConnectionString)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	ctx = privacy.DecisionContext(ctx, privacy.Allow)
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

	borik, err := client.Project.Create().
		SetName("Borik").
		SetDescription("a discord bot, written using discordgo, for ✨ breaking images ✨").
		SetStartDate(time.Date(2020, time.November, 5, 0, 0, 0, 0, time.UTC)).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test project: %v", err)
	}

	_, err = client.ProjectContributor.Create().
		SetRole(projectcontributor.RoleOwner).
		SetProject(borik).
		SetUser(riley).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test project contributor: %v", err)
	}

	_, err = client.ProjectContributor.Create().
		SetRole(projectcontributor.RoleOwner).
		SetProject(borik).
		SetUser(jack).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test project contributor: %v", err)
	}

	grackDb, err := client.Project.Create().
		SetName("GrackDB").
		SetStartDate(time.Date(2021, time.July, 30, 0, 0, 0, 0, time.UTC)).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test project: %v", err)
	}

	_, err = client.ProjectContributor.Create().
		SetRole(projectcontributor.RoleOwner).
		SetProject(grackDb).
		SetUser(riley).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test project contributor: %v", err)
	}

	_, err = client.ProjectContributor.Create().
		SetRole(projectcontributor.RoleOwner).
		SetProject(grackDb).
		SetUser(jack).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test project contributor: %v", err)
	}
}

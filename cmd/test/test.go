package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"github.com/fogo-sh/grackdb/ent"
	"github.com/fogo-sh/grackdb/ent/githubaccount"
	"github.com/fogo-sh/grackdb/ent/githuborganization"
	"github.com/fogo-sh/grackdb/ent/githuborganizationmember"
	"github.com/fogo-sh/grackdb/ent/user"
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

	testUser, err := client.User.Create().SetUsername("nint8835").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test user: %v", err)
		return
	}

	testDiscordAccount, err := client.DiscordAccount.Create().
		SetDiscordID("106162668032802816").
		SetUsername("nint8835").
		SetDiscriminator("0001").
		SetOwner(testUser).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test discord account: %v", err)
		return
	}

	fmt.Printf("Test discord account: %#+v\n", testDiscordAccount)

	testGithubAccount, err := client.GithubAccount.Create().
		SetUsername("nint8835").
		SetOwner(testUser).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test github account: %v", err)
		return
	}

	fmt.Printf("Test github account: %#+v\n", testGithubAccount)

	testGithubOrg, err := client.GithubOrganization.Create().
		SetName("fogo-sh").
		SetDisplayName("fogo.sh").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test github org: %v", err)
		return
	}

	fmt.Printf("Test github org: %#+v\n", testGithubOrg)

	testOrgMembership, err := client.GithubOrganizationMember.Create().
		SetRole(githuborganizationmember.RoleAdmin).
		SetOrganization(testGithubOrg).
		SetAccount(testGithubAccount).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating test github org membership: %v", err)
		return
	}

	fmt.Printf("Test github org membership: %#+v\n", testOrgMembership)

	testUsers, err := client.Debug().User.Query().
		Where(
			user.HasGithubAccountsWith(
				githubaccount.HasOrganizationMembershipsWith(
					githuborganizationmember.And(
						githuborganizationmember.RoleEQ(githuborganizationmember.RoleAdmin),
						githuborganizationmember.HasOrganizationWith(
							githuborganization.NameEQ("fogo-sh"),
						),
					),
				),
			),
		).
		All(ctx)
	if err != nil {
		log.Fatalf("failed querying test users: %v", err)
		return
	}
	fmt.Printf("Test query results: %#+v\n", testUsers)
}

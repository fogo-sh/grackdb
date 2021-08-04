package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/fogo-sh/grackdb/ent/privacy"
	_ "github.com/fogo-sh/grackdb/ent/runtime"
	"github.com/mitchellh/mapstructure"

	"entgo.io/ent/dialect"
	"github.com/fogo-sh/grackdb"
	"github.com/fogo-sh/grackdb/ent"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v2"
)

//go:embed testdata.yaml
var testdataYaml []byte

type TestItemType string

const (
	DiscordAccount           TestItemType = "DiscordAccount"
	DiscordBot               TestItemType = "DiscordBot"
	GithubAccount            TestItemType = "GithubAccount"
	GithubOrganization       TestItemType = "GithubOrganization"
	GithubOrganizationMember TestItemType = "GithubOrganizationMember"
	Project                  TestItemType = "Project"
	ProjectAssociation       TestItemType = "ProjectAssociation"
	ProjectContributor       TestItemType = "ProjectContributor"
	ProjectTechnology        TestItemType = "ProjectTechnology"
	Repository               TestItemType = "Repository"
	RepositoryTechnology     TestItemType = "RepositoryTechnology"
	Site                     TestItemType = "Site"
	Technology               TestItemType = "Technology"
	TechnologyAssociation    TestItemType = "TechnologyAssociation"
	User                     TestItemType = "User"
)

type TestItem struct {
	Name   string                 `yaml:"Name"`
	Type   TestItemType           `yaml:"Type"`
	Params map[string]interface{} `yaml:"Params"`
}

type TestData struct {
	Items []TestItem `yaml:"Items"`
}

func ToTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
		// Convert it by parsing
	}
}

func Decode(input map[string]interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}

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

	testdata := new(TestData)
	err = yaml.Unmarshal(testdataYaml, testdata)
	if err != nil {
		log.Fatalf("failed loading test data: %v", err)
	}

	for _, item := range testdata.Items {
		fmt.Printf("Creating %s %s\n", item.Type, item.Name)

		switch item.Type {
		case DiscordAccount:
			input := new(ent.CreateDiscordAccountInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding DiscordAccount: %v", err)
			}
			_, err = client.DiscordAccount.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating DiscordAccount: %v", err)
			}
		case DiscordBot:
			input := new(ent.CreateDiscordBotInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding DiscordBot: %v", err)
			}
			_, err = client.DiscordBot.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating DiscordBot: %v", err)
			}
		case GithubAccount:
			input := new(ent.CreateGithubAccountInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding GithubAccount: %v", err)
			}
			_, err = client.GithubAccount.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating GithubAccount: %v", err)
			}
		case GithubOrganization:
			input := new(ent.CreateGithubOrganizationInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding GithubOrganization: %v", err)
			}
			_, err = client.GithubOrganization.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating GithubOrganization: %v", err)
			}
		case GithubOrganizationMember:
			input := new(ent.CreateGithubOrganizationMemberInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding GithubOrganizationMember: %v", err)
			}
			_, err = client.GithubOrganizationMember.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating GithubOrganizationMember: %v", err)
			}
		case Project:
			input := new(ent.CreateProjectInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding Project: %v", err)
			}
			_, err = client.Project.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating Project: %v", err)
			}
		case ProjectAssociation:
			input := new(ent.CreateProjectAssociationInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding ProjectAssociation: %v", err)
			}
			_, err = client.ProjectAssociation.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating ProjectAssociation: %v", err)
			}
		case ProjectContributor:
			input := new(ent.CreateProjectContributorInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding ProjectContributor: %v", err)
			}
			_, err = client.ProjectContributor.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating ProjectContributor: %v", err)
			}
		case ProjectTechnology:
			input := new(ent.CreateProjectTechnologyInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding ProjectTechnology: %v", err)
			}
			_, err = client.ProjectTechnology.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating ProjectTechnology: %v", err)
			}
		case Repository:
			input := new(ent.CreateRepositoryInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding Repository: %v", err)
			}
			_, err = client.Repository.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating Repository: %v", err)
			}
		case RepositoryTechnology:
			input := new(ent.CreateRepositoryTechnologyInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding RepositoryTechnology: %v", err)
			}
			_, err = client.RepositoryTechnology.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating RepositoryTechnology: %v", err)
			}
		case Site:
			input := new(ent.CreateSiteInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding Site: %v", err)
			}
			_, err = client.Site.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating Site: %v", err)
			}
		case Technology:
			input := new(ent.CreateTechnologyInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding Technology: %v", err)
			}
			_, err = client.Technology.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating Technology: %v", err)
			}
		case TechnologyAssociation:
			input := new(ent.CreateTechnologyAssociationInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding TechnologyAssociation: %v", err)
			}
			_, err = client.TechnologyAssociation.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating TechnologyAssociation: %v", err)
			}
		case User:
			input := new(ent.CreateUserInput)
			err = Decode(item.Params, input)
			if err != nil {
				log.Fatalf("failed decoding User: %v", err)
			}
			_, err = client.User.Create().SetInput(*input).Save(ctx)
			if err != nil {
				log.Fatalf("failed creating User: %v", err)
			}
		default:
			log.Fatalf("unsupported resource type: %v", item.Type)
		}
	}

	fmt.Printf("Created %d test objects\n", len(testdata.Items))
}

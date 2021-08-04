// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/fogo-sh/grackdb/ent/githuborganizationmember"
	"github.com/fogo-sh/grackdb/ent/projectassociation"
	"github.com/fogo-sh/grackdb/ent/projectcontributor"
)

// CreateDiscordAccountInput represents a mutation input for creating discordaccounts.
type CreateDiscordAccountInput struct {
	DiscordID     string
	Username      string
	Discriminator string
	Owner         *int
	Bot           *int
}

// Mutate applies the CreateDiscordAccountInput on the DiscordAccountCreate builder.
func (i *CreateDiscordAccountInput) Mutate(m *DiscordAccountCreate) {
	m.SetDiscordID(i.DiscordID)
	m.SetUsername(i.Username)
	m.SetDiscriminator(i.Discriminator)
	if v := i.Owner; v != nil {
		m.SetOwnerID(*v)
	}
	if v := i.Bot; v != nil {
		m.SetBotID(*v)
	}
}

// SetInput applies the change-set in the CreateDiscordAccountInput on the create builder.
func (c *DiscordAccountCreate) SetInput(i CreateDiscordAccountInput) *DiscordAccountCreate {
	i.Mutate(c)
	return c
}

// UpdateDiscordAccountInput represents a mutation input for updating discordaccounts.
type UpdateDiscordAccountInput struct {
	Username      *string
	Discriminator *string
	Owner         *int
	ClearOwner    bool
	Bot           *int
	ClearBot      bool
}

// Mutate applies the UpdateDiscordAccountInput on the DiscordAccountMutation.
func (i *UpdateDiscordAccountInput) Mutate(m *DiscordAccountMutation) {
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if v := i.Discriminator; v != nil {
		m.SetDiscriminator(*v)
	}
	if i.ClearOwner {
		m.ClearOwner()
	}
	if v := i.Owner; v != nil {
		m.SetOwnerID(*v)
	}
	if i.ClearBot {
		m.ClearBot()
	}
	if v := i.Bot; v != nil {
		m.SetBotID(*v)
	}
}

// SetInput applies the change-set in the UpdateDiscordAccountInput on the update builder.
func (u *DiscordAccountUpdate) SetInput(i UpdateDiscordAccountInput) *DiscordAccountUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateDiscordAccountInput on the update-one builder.
func (u *DiscordAccountUpdateOne) SetInput(i UpdateDiscordAccountInput) *DiscordAccountUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateDiscordBotInput represents a mutation input for creating discordbots.
type CreateDiscordBotInput struct {
	Account    int
	Project    int
	Repository *int
}

// Mutate applies the CreateDiscordBotInput on the DiscordBotCreate builder.
func (i *CreateDiscordBotInput) Mutate(m *DiscordBotCreate) {
	m.SetAccountID(i.Account)
	m.SetProjectID(i.Project)
	if v := i.Repository; v != nil {
		m.SetRepositoryID(*v)
	}
}

// SetInput applies the change-set in the CreateDiscordBotInput on the create builder.
func (c *DiscordBotCreate) SetInput(i CreateDiscordBotInput) *DiscordBotCreate {
	i.Mutate(c)
	return c
}

// UpdateDiscordBotInput represents a mutation input for updating discordbots.
type UpdateDiscordBotInput struct {
	Account         *int
	ClearAccount    bool
	Project         *int
	ClearProject    bool
	Repository      *int
	ClearRepository bool
}

// Mutate applies the UpdateDiscordBotInput on the DiscordBotMutation.
func (i *UpdateDiscordBotInput) Mutate(m *DiscordBotMutation) {
	if i.ClearAccount {
		m.ClearAccount()
	}
	if v := i.Account; v != nil {
		m.SetAccountID(*v)
	}
	if i.ClearProject {
		m.ClearProject()
	}
	if v := i.Project; v != nil {
		m.SetProjectID(*v)
	}
	if i.ClearRepository {
		m.ClearRepository()
	}
	if v := i.Repository; v != nil {
		m.SetRepositoryID(*v)
	}
}

// SetInput applies the change-set in the UpdateDiscordBotInput on the update builder.
func (u *DiscordBotUpdate) SetInput(i UpdateDiscordBotInput) *DiscordBotUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateDiscordBotInput on the update-one builder.
func (u *DiscordBotUpdateOne) SetInput(i UpdateDiscordBotInput) *DiscordBotUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateGithubAccountInput represents a mutation input for creating githubaccounts.
type CreateGithubAccountInput struct {
	Username                string
	Owner                   int
	OrganizationMemberships []int
	Repositories            []int
}

// Mutate applies the CreateGithubAccountInput on the GithubAccountCreate builder.
func (i *CreateGithubAccountInput) Mutate(m *GithubAccountCreate) {
	m.SetUsername(i.Username)
	m.SetOwnerID(i.Owner)
	if ids := i.OrganizationMemberships; len(ids) > 0 {
		m.AddOrganizationMembershipIDs(ids...)
	}
	if ids := i.Repositories; len(ids) > 0 {
		m.AddRepositoryIDs(ids...)
	}
}

// SetInput applies the change-set in the CreateGithubAccountInput on the create builder.
func (c *GithubAccountCreate) SetInput(i CreateGithubAccountInput) *GithubAccountCreate {
	i.Mutate(c)
	return c
}

// UpdateGithubAccountInput represents a mutation input for updating githubaccounts.
type UpdateGithubAccountInput struct {
	Username                        *string
	Owner                           *int
	ClearOwner                      bool
	AddOrganizationMembershipIDs    []int
	RemoveOrganizationMembershipIDs []int
	AddRepositoryIDs                []int
	RemoveRepositoryIDs             []int
}

// Mutate applies the UpdateGithubAccountInput on the GithubAccountMutation.
func (i *UpdateGithubAccountInput) Mutate(m *GithubAccountMutation) {
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if i.ClearOwner {
		m.ClearOwner()
	}
	if v := i.Owner; v != nil {
		m.SetOwnerID(*v)
	}
	if ids := i.AddOrganizationMembershipIDs; len(ids) > 0 {
		m.AddOrganizationMembershipIDs(ids...)
	}
	if ids := i.RemoveOrganizationMembershipIDs; len(ids) > 0 {
		m.RemoveOrganizationMembershipIDs(ids...)
	}
	if ids := i.AddRepositoryIDs; len(ids) > 0 {
		m.AddRepositoryIDs(ids...)
	}
	if ids := i.RemoveRepositoryIDs; len(ids) > 0 {
		m.RemoveRepositoryIDs(ids...)
	}
}

// SetInput applies the change-set in the UpdateGithubAccountInput on the update builder.
func (u *GithubAccountUpdate) SetInput(i UpdateGithubAccountInput) *GithubAccountUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateGithubAccountInput on the update-one builder.
func (u *GithubAccountUpdateOne) SetInput(i UpdateGithubAccountInput) *GithubAccountUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateGithubOrganizationInput represents a mutation input for creating githuborganizations.
type CreateGithubOrganizationInput struct {
	Name         string
	DisplayName  *string
	Members      []int
	Repositories []int
}

// Mutate applies the CreateGithubOrganizationInput on the GithubOrganizationCreate builder.
func (i *CreateGithubOrganizationInput) Mutate(m *GithubOrganizationCreate) {
	m.SetName(i.Name)
	if v := i.DisplayName; v != nil {
		m.SetDisplayName(*v)
	}
	if ids := i.Members; len(ids) > 0 {
		m.AddMemberIDs(ids...)
	}
	if ids := i.Repositories; len(ids) > 0 {
		m.AddRepositoryIDs(ids...)
	}
}

// SetInput applies the change-set in the CreateGithubOrganizationInput on the create builder.
func (c *GithubOrganizationCreate) SetInput(i CreateGithubOrganizationInput) *GithubOrganizationCreate {
	i.Mutate(c)
	return c
}

// UpdateGithubOrganizationInput represents a mutation input for updating githuborganizations.
type UpdateGithubOrganizationInput struct {
	Name                *string
	DisplayName         *string
	ClearDisplayName    bool
	AddMemberIDs        []int
	RemoveMemberIDs     []int
	AddRepositoryIDs    []int
	RemoveRepositoryIDs []int
}

// Mutate applies the UpdateGithubOrganizationInput on the GithubOrganizationMutation.
func (i *UpdateGithubOrganizationInput) Mutate(m *GithubOrganizationMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearDisplayName {
		m.ClearDisplayName()
	}
	if v := i.DisplayName; v != nil {
		m.SetDisplayName(*v)
	}
	if ids := i.AddMemberIDs; len(ids) > 0 {
		m.AddMemberIDs(ids...)
	}
	if ids := i.RemoveMemberIDs; len(ids) > 0 {
		m.RemoveMemberIDs(ids...)
	}
	if ids := i.AddRepositoryIDs; len(ids) > 0 {
		m.AddRepositoryIDs(ids...)
	}
	if ids := i.RemoveRepositoryIDs; len(ids) > 0 {
		m.RemoveRepositoryIDs(ids...)
	}
}

// SetInput applies the change-set in the UpdateGithubOrganizationInput on the update builder.
func (u *GithubOrganizationUpdate) SetInput(i UpdateGithubOrganizationInput) *GithubOrganizationUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateGithubOrganizationInput on the update-one builder.
func (u *GithubOrganizationUpdateOne) SetInput(i UpdateGithubOrganizationInput) *GithubOrganizationUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateGithubOrganizationMemberInput represents a mutation input for creating githuborganizationmembers.
type CreateGithubOrganizationMemberInput struct {
	Role         *githuborganizationmember.Role
	Organization int
	Account      int
}

// Mutate applies the CreateGithubOrganizationMemberInput on the GithubOrganizationMemberCreate builder.
func (i *CreateGithubOrganizationMemberInput) Mutate(m *GithubOrganizationMemberCreate) {
	if v := i.Role; v != nil {
		m.SetRole(*v)
	}
	m.SetOrganizationID(i.Organization)
	m.SetAccountID(i.Account)
}

// SetInput applies the change-set in the CreateGithubOrganizationMemberInput on the create builder.
func (c *GithubOrganizationMemberCreate) SetInput(i CreateGithubOrganizationMemberInput) *GithubOrganizationMemberCreate {
	i.Mutate(c)
	return c
}

// UpdateGithubOrganizationMemberInput represents a mutation input for updating githuborganizationmembers.
type UpdateGithubOrganizationMemberInput struct {
	Role              *githuborganizationmember.Role
	Organization      *int
	ClearOrganization bool
	Account           *int
	ClearAccount      bool
}

// Mutate applies the UpdateGithubOrganizationMemberInput on the GithubOrganizationMemberMutation.
func (i *UpdateGithubOrganizationMemberInput) Mutate(m *GithubOrganizationMemberMutation) {
	if v := i.Role; v != nil {
		m.SetRole(*v)
	}
	if i.ClearOrganization {
		m.ClearOrganization()
	}
	if v := i.Organization; v != nil {
		m.SetOrganizationID(*v)
	}
	if i.ClearAccount {
		m.ClearAccount()
	}
	if v := i.Account; v != nil {
		m.SetAccountID(*v)
	}
}

// SetInput applies the change-set in the UpdateGithubOrganizationMemberInput on the update builder.
func (u *GithubOrganizationMemberUpdate) SetInput(i UpdateGithubOrganizationMemberInput) *GithubOrganizationMemberUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateGithubOrganizationMemberInput on the update-one builder.
func (u *GithubOrganizationMemberUpdateOne) SetInput(i UpdateGithubOrganizationMemberInput) *GithubOrganizationMemberUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateProjectInput represents a mutation input for creating projects.
type CreateProjectInput struct {
	Name           string
	Description    *string
	StartDate      time.Time
	EndDate        *time.Time
	Contributors   []int
	ParentProjects []int
	ChildProjects  []int
	Repositories   []int
	DiscordBots    []int
	Sites          []int
}

// Mutate applies the CreateProjectInput on the ProjectCreate builder.
func (i *CreateProjectInput) Mutate(m *ProjectCreate) {
	m.SetName(i.Name)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetStartDate(i.StartDate)
	if v := i.EndDate; v != nil {
		m.SetEndDate(*v)
	}
	if ids := i.Contributors; len(ids) > 0 {
		m.AddContributorIDs(ids...)
	}
	if ids := i.ParentProjects; len(ids) > 0 {
		m.AddParentProjectIDs(ids...)
	}
	if ids := i.ChildProjects; len(ids) > 0 {
		m.AddChildProjectIDs(ids...)
	}
	if ids := i.Repositories; len(ids) > 0 {
		m.AddRepositoryIDs(ids...)
	}
	if ids := i.DiscordBots; len(ids) > 0 {
		m.AddDiscordBotIDs(ids...)
	}
	if ids := i.Sites; len(ids) > 0 {
		m.AddSiteIDs(ids...)
	}
}

// SetInput applies the change-set in the CreateProjectInput on the create builder.
func (c *ProjectCreate) SetInput(i CreateProjectInput) *ProjectCreate {
	i.Mutate(c)
	return c
}

// UpdateProjectInput represents a mutation input for updating projects.
type UpdateProjectInput struct {
	Name                   *string
	Description            *string
	ClearDescription       bool
	StartDate              *time.Time
	EndDate                *time.Time
	ClearEndDate           bool
	AddContributorIDs      []int
	RemoveContributorIDs   []int
	AddParentProjectIDs    []int
	RemoveParentProjectIDs []int
	AddChildProjectIDs     []int
	RemoveChildProjectIDs  []int
	AddRepositoryIDs       []int
	RemoveRepositoryIDs    []int
	AddDiscordBotIDs       []int
	RemoveDiscordBotIDs    []int
	AddSiteIDs             []int
	RemoveSiteIDs          []int
}

// Mutate applies the UpdateProjectInput on the ProjectMutation.
func (i *UpdateProjectInput) Mutate(m *ProjectMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearDescription {
		m.ClearDescription()
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.StartDate; v != nil {
		m.SetStartDate(*v)
	}
	if i.ClearEndDate {
		m.ClearEndDate()
	}
	if v := i.EndDate; v != nil {
		m.SetEndDate(*v)
	}
	if ids := i.AddContributorIDs; len(ids) > 0 {
		m.AddContributorIDs(ids...)
	}
	if ids := i.RemoveContributorIDs; len(ids) > 0 {
		m.RemoveContributorIDs(ids...)
	}
	if ids := i.AddParentProjectIDs; len(ids) > 0 {
		m.AddParentProjectIDs(ids...)
	}
	if ids := i.RemoveParentProjectIDs; len(ids) > 0 {
		m.RemoveParentProjectIDs(ids...)
	}
	if ids := i.AddChildProjectIDs; len(ids) > 0 {
		m.AddChildProjectIDs(ids...)
	}
	if ids := i.RemoveChildProjectIDs; len(ids) > 0 {
		m.RemoveChildProjectIDs(ids...)
	}
	if ids := i.AddRepositoryIDs; len(ids) > 0 {
		m.AddRepositoryIDs(ids...)
	}
	if ids := i.RemoveRepositoryIDs; len(ids) > 0 {
		m.RemoveRepositoryIDs(ids...)
	}
	if ids := i.AddDiscordBotIDs; len(ids) > 0 {
		m.AddDiscordBotIDs(ids...)
	}
	if ids := i.RemoveDiscordBotIDs; len(ids) > 0 {
		m.RemoveDiscordBotIDs(ids...)
	}
	if ids := i.AddSiteIDs; len(ids) > 0 {
		m.AddSiteIDs(ids...)
	}
	if ids := i.RemoveSiteIDs; len(ids) > 0 {
		m.RemoveSiteIDs(ids...)
	}
}

// SetInput applies the change-set in the UpdateProjectInput on the update builder.
func (u *ProjectUpdate) SetInput(i UpdateProjectInput) *ProjectUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateProjectInput on the update-one builder.
func (u *ProjectUpdateOne) SetInput(i UpdateProjectInput) *ProjectUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateProjectAssociationInput represents a mutation input for creating projectassociations.
type CreateProjectAssociationInput struct {
	Type   projectassociation.Type
	Parent int
	Child  int
}

// Mutate applies the CreateProjectAssociationInput on the ProjectAssociationCreate builder.
func (i *CreateProjectAssociationInput) Mutate(m *ProjectAssociationCreate) {
	m.SetType(i.Type)
	m.SetParentID(i.Parent)
	m.SetChildID(i.Child)
}

// SetInput applies the change-set in the CreateProjectAssociationInput on the create builder.
func (c *ProjectAssociationCreate) SetInput(i CreateProjectAssociationInput) *ProjectAssociationCreate {
	i.Mutate(c)
	return c
}

// UpdateProjectAssociationInput represents a mutation input for updating projectassociations.
type UpdateProjectAssociationInput struct {
	Type        *projectassociation.Type
	Parent      *int
	ClearParent bool
	Child       *int
	ClearChild  bool
}

// Mutate applies the UpdateProjectAssociationInput on the ProjectAssociationMutation.
func (i *UpdateProjectAssociationInput) Mutate(m *ProjectAssociationMutation) {
	if v := i.Type; v != nil {
		m.SetType(*v)
	}
	if i.ClearParent {
		m.ClearParent()
	}
	if v := i.Parent; v != nil {
		m.SetParentID(*v)
	}
	if i.ClearChild {
		m.ClearChild()
	}
	if v := i.Child; v != nil {
		m.SetChildID(*v)
	}
}

// SetInput applies the change-set in the UpdateProjectAssociationInput on the update builder.
func (u *ProjectAssociationUpdate) SetInput(i UpdateProjectAssociationInput) *ProjectAssociationUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateProjectAssociationInput on the update-one builder.
func (u *ProjectAssociationUpdateOne) SetInput(i UpdateProjectAssociationInput) *ProjectAssociationUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateProjectContributorInput represents a mutation input for creating projectcontributors.
type CreateProjectContributorInput struct {
	Role    projectcontributor.Role
	Project int
	User    int
}

// Mutate applies the CreateProjectContributorInput on the ProjectContributorCreate builder.
func (i *CreateProjectContributorInput) Mutate(m *ProjectContributorCreate) {
	m.SetRole(i.Role)
	m.SetProjectID(i.Project)
	m.SetUserID(i.User)
}

// SetInput applies the change-set in the CreateProjectContributorInput on the create builder.
func (c *ProjectContributorCreate) SetInput(i CreateProjectContributorInput) *ProjectContributorCreate {
	i.Mutate(c)
	return c
}

// UpdateProjectContributorInput represents a mutation input for updating projectcontributors.
type UpdateProjectContributorInput struct {
	Role         *projectcontributor.Role
	Project      *int
	ClearProject bool
	User         *int
	ClearUser    bool
}

// Mutate applies the UpdateProjectContributorInput on the ProjectContributorMutation.
func (i *UpdateProjectContributorInput) Mutate(m *ProjectContributorMutation) {
	if v := i.Role; v != nil {
		m.SetRole(*v)
	}
	if i.ClearProject {
		m.ClearProject()
	}
	if v := i.Project; v != nil {
		m.SetProjectID(*v)
	}
	if i.ClearUser {
		m.ClearUser()
	}
	if v := i.User; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the UpdateProjectContributorInput on the update builder.
func (u *ProjectContributorUpdate) SetInput(i UpdateProjectContributorInput) *ProjectContributorUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateProjectContributorInput on the update-one builder.
func (u *ProjectContributorUpdateOne) SetInput(i UpdateProjectContributorInput) *ProjectContributorUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateRepositoryInput represents a mutation input for creating repositories.
type CreateRepositoryInput struct {
	Name               string
	Description        *string
	Project            int
	GithubAccount      *int
	GithubOrganization *int
	DiscordBots        []int
	Sites              []int
}

// Mutate applies the CreateRepositoryInput on the RepositoryCreate builder.
func (i *CreateRepositoryInput) Mutate(m *RepositoryCreate) {
	m.SetName(i.Name)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetProjectID(i.Project)
	if v := i.GithubAccount; v != nil {
		m.SetGithubAccountID(*v)
	}
	if v := i.GithubOrganization; v != nil {
		m.SetGithubOrganizationID(*v)
	}
	if ids := i.DiscordBots; len(ids) > 0 {
		m.AddDiscordBotIDs(ids...)
	}
	if ids := i.Sites; len(ids) > 0 {
		m.AddSiteIDs(ids...)
	}
}

// SetInput applies the change-set in the CreateRepositoryInput on the create builder.
func (c *RepositoryCreate) SetInput(i CreateRepositoryInput) *RepositoryCreate {
	i.Mutate(c)
	return c
}

// UpdateRepositoryInput represents a mutation input for updating repositories.
type UpdateRepositoryInput struct {
	Name                    *string
	Description             *string
	ClearDescription        bool
	Project                 *int
	ClearProject            bool
	GithubAccount           *int
	ClearGithubAccount      bool
	GithubOrganization      *int
	ClearGithubOrganization bool
	AddDiscordBotIDs        []int
	RemoveDiscordBotIDs     []int
	AddSiteIDs              []int
	RemoveSiteIDs           []int
}

// Mutate applies the UpdateRepositoryInput on the RepositoryMutation.
func (i *UpdateRepositoryInput) Mutate(m *RepositoryMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearDescription {
		m.ClearDescription()
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if i.ClearProject {
		m.ClearProject()
	}
	if v := i.Project; v != nil {
		m.SetProjectID(*v)
	}
	if i.ClearGithubAccount {
		m.ClearGithubAccount()
	}
	if v := i.GithubAccount; v != nil {
		m.SetGithubAccountID(*v)
	}
	if i.ClearGithubOrganization {
		m.ClearGithubOrganization()
	}
	if v := i.GithubOrganization; v != nil {
		m.SetGithubOrganizationID(*v)
	}
	if ids := i.AddDiscordBotIDs; len(ids) > 0 {
		m.AddDiscordBotIDs(ids...)
	}
	if ids := i.RemoveDiscordBotIDs; len(ids) > 0 {
		m.RemoveDiscordBotIDs(ids...)
	}
	if ids := i.AddSiteIDs; len(ids) > 0 {
		m.AddSiteIDs(ids...)
	}
	if ids := i.RemoveSiteIDs; len(ids) > 0 {
		m.RemoveSiteIDs(ids...)
	}
}

// SetInput applies the change-set in the UpdateRepositoryInput on the update builder.
func (u *RepositoryUpdate) SetInput(i UpdateRepositoryInput) *RepositoryUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateRepositoryInput on the update-one builder.
func (u *RepositoryUpdateOne) SetInput(i UpdateRepositoryInput) *RepositoryUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateSiteInput represents a mutation input for creating sites.
type CreateSiteInput struct {
	URL        string
	Project    int
	Repository *int
}

// Mutate applies the CreateSiteInput on the SiteCreate builder.
func (i *CreateSiteInput) Mutate(m *SiteCreate) {
	m.SetURL(i.URL)
	m.SetProjectID(i.Project)
	if v := i.Repository; v != nil {
		m.SetRepositoryID(*v)
	}
}

// SetInput applies the change-set in the CreateSiteInput on the create builder.
func (c *SiteCreate) SetInput(i CreateSiteInput) *SiteCreate {
	i.Mutate(c)
	return c
}

// UpdateSiteInput represents a mutation input for updating sites.
type UpdateSiteInput struct {
	URL             *string
	Project         *int
	ClearProject    bool
	Repository      *int
	ClearRepository bool
}

// Mutate applies the UpdateSiteInput on the SiteMutation.
func (i *UpdateSiteInput) Mutate(m *SiteMutation) {
	if v := i.URL; v != nil {
		m.SetURL(*v)
	}
	if i.ClearProject {
		m.ClearProject()
	}
	if v := i.Project; v != nil {
		m.SetProjectID(*v)
	}
	if i.ClearRepository {
		m.ClearRepository()
	}
	if v := i.Repository; v != nil {
		m.SetRepositoryID(*v)
	}
}

// SetInput applies the change-set in the UpdateSiteInput on the update builder.
func (u *SiteUpdate) SetInput(i UpdateSiteInput) *SiteUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSiteInput on the update-one builder.
func (u *SiteUpdateOne) SetInput(i UpdateSiteInput) *SiteUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Username             string
	AvatarURL            *string
	DiscordAccounts      []int
	GithubAccounts       []int
	ProjectContributions []int
}

// Mutate applies the CreateUserInput on the UserCreate builder.
func (i *CreateUserInput) Mutate(m *UserCreate) {
	m.SetUsername(i.Username)
	if v := i.AvatarURL; v != nil {
		m.SetAvatarURL(*v)
	}
	if ids := i.DiscordAccounts; len(ids) > 0 {
		m.AddDiscordAccountIDs(ids...)
	}
	if ids := i.GithubAccounts; len(ids) > 0 {
		m.AddGithubAccountIDs(ids...)
	}
	if ids := i.ProjectContributions; len(ids) > 0 {
		m.AddProjectContributionIDs(ids...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the create builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c)
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Username                     *string
	AvatarURL                    *string
	ClearAvatarURL               bool
	AddDiscordAccountIDs         []int
	RemoveDiscordAccountIDs      []int
	AddGithubAccountIDs          []int
	RemoveGithubAccountIDs       []int
	AddProjectContributionIDs    []int
	RemoveProjectContributionIDs []int
}

// Mutate applies the UpdateUserInput on the UserMutation.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if i.ClearAvatarURL {
		m.ClearAvatarURL()
	}
	if v := i.AvatarURL; v != nil {
		m.SetAvatarURL(*v)
	}
	if ids := i.AddDiscordAccountIDs; len(ids) > 0 {
		m.AddDiscordAccountIDs(ids...)
	}
	if ids := i.RemoveDiscordAccountIDs; len(ids) > 0 {
		m.RemoveDiscordAccountIDs(ids...)
	}
	if ids := i.AddGithubAccountIDs; len(ids) > 0 {
		m.AddGithubAccountIDs(ids...)
	}
	if ids := i.RemoveGithubAccountIDs; len(ids) > 0 {
		m.RemoveGithubAccountIDs(ids...)
	}
	if ids := i.AddProjectContributionIDs; len(ids) > 0 {
		m.AddProjectContributionIDs(ids...)
	}
	if ids := i.RemoveProjectContributionIDs; len(ids) > 0 {
		m.RemoveProjectContributionIDs(ids...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the update builder.
func (u *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateUserInput on the update-one builder.
func (u *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

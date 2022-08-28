// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/fogo-sh/grackdb/ent/schema","Package":"github.com/fogo-sh/grackdb/ent","Schemas":[{"name":"DiscordAccount","config":{"Table":""},"edges":[{"name":"owner","type":"User","ref_name":"discord_accounts","unique":true,"inverse":true,"annotations":{"EntGQL":{}}},{"name":"bot","type":"DiscordBot","ref_name":"account","unique":true,"inverse":true,"annotations":{"EntGQL":{}}}],"fields":[{"name":"discord_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"immutable":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"DISCORD_ID"}}},{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"USERNAME"}}},{"name":"discriminator","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":4,"validators":2,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"DISCRIMINATOR"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"discordAccounts"},"RelayConnection":true}}},{"name":"DiscordBot","config":{"Table":""},"edges":[{"name":"account","type":"DiscordAccount","unique":true,"required":true,"annotations":{"EntGQL":{}}},{"name":"project","type":"Project","ref_name":"discord_bots","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}},{"name":"repository","type":"Repository","ref_name":"discord_bots","unique":true,"inverse":true,"annotations":{"EntGQL":{}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"discordBots"},"RelayConnection":true}}},{"name":"GithubAccount","config":{"Table":""},"edges":[{"name":"owner","type":"User","ref_name":"github_accounts","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}},{"name":"organization_memberships","type":"GithubOrganizationMember","annotations":{"EntGQL":{"Mapping":["organizationMemberships"],"Unbind":true}}},{"name":"repositories","type":"Repository","annotations":{"EntGQL":{}}}],"fields":[{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"USERNAME"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"githubAccounts"},"RelayConnection":true}}},{"name":"GithubOrganization","config":{"Table":""},"edges":[{"name":"members","type":"GithubOrganizationMember","annotations":{"EntGQL":{}}},{"name":"repositories","type":"Repository","annotations":{"EntGQL":{}}}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}},{"name":"display_name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"DISPLAY_NAME"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"githubOrganizations"},"RelayConnection":true}}},{"name":"GithubOrganizationMember","config":{"Table":""},"edges":[{"name":"organization","type":"GithubOrganization","ref_name":"members","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}},{"name":"account","type":"GithubAccount","ref_name":"organization_memberships","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}}],"fields":[{"name":"role","type":{"Type":6,"Ident":"githuborganizationmember.Role","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"Admin","V":"ADMIN"},{"N":"Member","V":"MEMBER"}],"default":true,"default_value":"MEMBER","default_kind":24,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"ROLE"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"githubOrganizationMembers"},"RelayConnection":true}}},{"name":"Project","config":{"Table":""},"edges":[{"name":"contributors","type":"ProjectContributor","annotations":{"EntGQL":{}}},{"name":"parent_projects","type":"ProjectAssociation","annotations":{"EntGQL":{"Mapping":["parentProjects"],"Unbind":true}}},{"name":"child_projects","type":"ProjectAssociation","annotations":{"EntGQL":{"Mapping":["childProjects"],"Unbind":true}}},{"name":"repositories","type":"Repository","annotations":{"EntGQL":{}}},{"name":"discord_bots","type":"DiscordBot","annotations":{"EntGQL":{"Mapping":["discordBots"],"Unbind":true}}},{"name":"sites","type":"Site","annotations":{"EntGQL":{}}},{"name":"technologies","type":"ProjectTechnology","annotations":{"EntGQL":{}}}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"DESCRIPTION"}}},{"name":"start_date","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"START_DATE"}}},{"name":"end_date","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"END_DATE"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"projects"},"RelayConnection":true}}},{"name":"ProjectAssociation","config":{"Table":""},"edges":[{"name":"parent","type":"Project","ref_name":"child_projects","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}},{"name":"child","type":"Project","ref_name":"parent_projects","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}}],"fields":[{"name":"type","type":{"Type":6,"Ident":"projectassociation.Type","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"BasedOff","V":"BASED_OFF"},{"N":"Replaces","V":"REPLACES"},{"N":"InspiredBy","V":"INSPIRED_BY"},{"N":"Related","V":"RELATED"}],"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"TYPE"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"projectAssociations"},"RelayConnection":true}}},{"name":"ProjectContributor","config":{"Table":""},"edges":[{"name":"project","type":"Project","ref_name":"contributors","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}},{"name":"user","type":"User","ref_name":"project_contributions","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}}],"fields":[{"name":"role","type":{"Type":6,"Ident":"projectcontributor.Role","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"Owner","V":"OWNER"},{"N":"Contributor","V":"CONTRIBUTOR"}],"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"ROLE"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"projectContributors"},"RelayConnection":true}}},{"name":"ProjectTechnology","config":{"Table":""},"edges":[{"name":"project","type":"Project","ref_name":"technologies","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}},{"name":"technology","type":"Technology","ref_name":"projects","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}}],"fields":[{"name":"type","type":{"Type":6,"Ident":"projecttechnology.Type","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"WrittenIn","V":"WRITTEN_IN"},{"N":"Implements","V":"IMPLEMENTS"},{"N":"Uses","V":"USES"},{"N":"Contains","V":"CONTAINS"}],"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"TYPE"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"projectTechnologies"},"RelayConnection":true}}},{"name":"Repository","config":{"Table":""},"edges":[{"name":"project","type":"Project","ref_name":"repositories","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}},{"name":"github_account","type":"GithubAccount","ref_name":"repositories","unique":true,"inverse":true,"annotations":{"EntGQL":{"Mapping":["githubAccount"],"Unbind":true}}},{"name":"github_organization","type":"GithubOrganization","ref_name":"repositories","unique":true,"inverse":true,"annotations":{"EntGQL":{"Mapping":["githubOrganization"],"Unbind":true}}},{"name":"discord_bots","type":"DiscordBot","annotations":{"EntGQL":{"Mapping":["discordBots"],"Unbind":true}}},{"name":"sites","type":"Site","annotations":{"EntGQL":{}}},{"name":"technologies","type":"RepositoryTechnology","annotations":{"EntGQL":{}}}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"DESCRIPTION"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"repositories"},"RelayConnection":true}}},{"name":"RepositoryTechnology","config":{"Table":""},"edges":[{"name":"repository","type":"Repository","ref_name":"technologies","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}},{"name":"technology","type":"Technology","ref_name":"repositories","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}}],"fields":[{"name":"type","type":{"Type":6,"Ident":"repositorytechnology.Type","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"WrittenIn","V":"WRITTEN_IN"},{"N":"Implements","V":"IMPLEMENTS"},{"N":"Uses","V":"USES"},{"N":"Contains","V":"CONTAINS"}],"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"TYPE"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"repositoryTechnologies"},"RelayConnection":true}}},{"name":"Site","config":{"Table":""},"edges":[{"name":"project","type":"Project","ref_name":"sites","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}},{"name":"repository","type":"Repository","ref_name":"sites","unique":true,"inverse":true,"annotations":{"EntGQL":{}}}],"fields":[{"name":"url","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"URL"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"sites"},"RelayConnection":true}}},{"name":"Technology","config":{"Table":""},"edges":[{"name":"parent_technologies","type":"TechnologyAssociation","annotations":{"EntGQL":{"Mapping":["parentTechnologies"],"Unbind":true}}},{"name":"child_technologies","type":"TechnologyAssociation","annotations":{"EntGQL":{"Mapping":["childTechnologies"],"Unbind":true}}},{"name":"projects","type":"ProjectTechnology","annotations":{"EntGQL":{}}},{"name":"repositories","type":"RepositoryTechnology","annotations":{"EntGQL":{}}}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"DESCRIPTION"}}},{"name":"colour","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"COLOUR"}}},{"name":"type","type":{"Type":6,"Ident":"technology.Type","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"Library","V":"LIBRARY"},{"N":"Language","V":"LANGUAGE"},{"N":"Algorithm","V":"ALGORITHM"},{"N":"Database","V":"DATABASE"},{"N":"Protocol","V":"PROTOCOL"},{"N":"Service","V":"SERVICE"}],"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"TYPE"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"technologies"},"RelayConnection":true}}},{"name":"TechnologyAssociation","config":{"Table":""},"edges":[{"name":"parent","type":"Technology","ref_name":"child_technologies","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}},{"name":"child","type":"Technology","ref_name":"parent_technologies","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{}}}],"fields":[{"name":"type","type":{"Type":6,"Ident":"technologyassociation.Type","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"WrittenIn","V":"WRITTEN_IN"},{"N":"Implements","V":"IMPLEMENTS"},{"N":"Uses","V":"USES"}],"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"TYPE"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"technologyAssociations"},"RelayConnection":true}}},{"name":"User","config":{"Table":""},"edges":[{"name":"discord_accounts","type":"DiscordAccount","annotations":{"EntGQL":{"Mapping":["discordAccounts"],"Unbind":true}}},{"name":"github_accounts","type":"GithubAccount","annotations":{"EntGQL":{"Mapping":["githubAccounts"],"Unbind":true}}},{"name":"project_contributions","type":"ProjectContributor","annotations":{"EntGQL":{"Mapping":["projectContributions"],"Unbind":true}}}],"fields":[{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"USERNAME"}}},{"name":"avatar_url","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"Mapping":["avatarUrl"],"Unbind":true}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{"Name":"users"},"RelayConnection":true}}}],"Features":["namedges","privacy","schema/snapshot"]}`

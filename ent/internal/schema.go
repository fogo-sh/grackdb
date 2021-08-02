// Code generated by entc, DO NOT EDIT.

// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/fogo-sh/grackdb/ent/schema","Package":"github.com/fogo-sh/grackdb/ent","Schemas":[{"name":"DiscordAccount","config":{"Table":""},"edges":[{"name":"owner","type":"User","ref_name":"discord_accounts","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"discord_id","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"unique":true,"immutable":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"DISCORD_ID"}}},{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"USERNAME"}}},{"name":"discriminator","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":4,"validators":2,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"DISCRIMINATOR"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"GithubAccount","config":{"Table":""},"edges":[{"name":"owner","type":"User","ref_name":"github_accounts","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{"Bind":true}}},{"name":"organization_memberships","type":"GithubOrganizationMember","annotations":{"EntGQL":{"Mapping":["organizationMemberships"]}}},{"name":"repositories","type":"Repository","annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"unique":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"USERNAME"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"GithubOrganization","config":{"Table":""},"edges":[{"name":"members","type":"GithubOrganizationMember","annotations":{"EntGQL":{"Bind":true}}},{"name":"repositories","type":"Repository","annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"unique":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}},{"name":"display_name","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"DISPLAY_NAME"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"GithubOrganizationMember","config":{"Table":""},"edges":[{"name":"organization","type":"GithubOrganization","ref_name":"members","unique":true,"inverse":true,"annotations":{"EntGQL":{"Bind":true}}},{"name":"account","type":"GithubAccount","ref_name":"organization_memberships","unique":true,"inverse":true,"annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"role","type":{"Type":6,"Ident":"githuborganizationmember.Role","PkgPath":"","Nillable":false,"RType":null},"enums":[{"N":"Admin","V":"ADMIN"},{"N":"Member","V":"MEMBER"}],"default":true,"default_value":"MEMBER","default_kind":24,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"ROLE"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"Project","config":{"Table":""},"edges":[{"name":"contributors","type":"ProjectContributor","annotations":{"EntGQL":{"Bind":true}}},{"name":"parent_projects","type":"ProjectAssociation","annotations":{"EntGQL":{"Mapping":["parentProjects"]}}},{"name":"child_projects","type":"ProjectAssociation","annotations":{"EntGQL":{"Mapping":["childProjects"]}}},{"name":"repositories","type":"Repository","annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"DESCRIPTION"}}},{"name":"start_date","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"START_DATE"}}},{"name":"end_date","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"END_DATE"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"ProjectAssociation","config":{"Table":""},"edges":[{"name":"parent","type":"Project","ref_name":"child_projects","unique":true,"inverse":true,"annotations":{"EntGQL":{"Bind":true}}},{"name":"child","type":"Project","ref_name":"parent_projects","unique":true,"inverse":true,"annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"type","type":{"Type":6,"Ident":"projectassociation.Type","PkgPath":"","Nillable":false,"RType":null},"enums":[{"N":"BasedOff","V":"BASED_OFF"},{"N":"Replaces","V":"REPLACES"},{"N":"InspiredBy","V":"INSPIRED_BY"},{"N":"Related","V":"RELATED"}],"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"TYPE"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"ProjectContributor","config":{"Table":""},"edges":[{"name":"project","type":"Project","ref_name":"contributors","unique":true,"inverse":true,"annotations":{"EntGQL":{"Bind":true}}},{"name":"user","type":"User","ref_name":"project_contributions","unique":true,"inverse":true,"annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"role","type":{"Type":6,"Ident":"projectcontributor.Role","PkgPath":"","Nillable":false,"RType":null},"enums":[{"N":"Owner","V":"OWNER"},{"N":"Contributor","V":"CONTRIBUTOR"}],"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"ROLE"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"Repository","config":{"Table":""},"edges":[{"name":"project","type":"Project","ref_name":"repositories","unique":true,"inverse":true,"annotations":{"EntGQL":{"Bind":true}}},{"name":"github_account","type":"GithubAccount","ref_name":"repositories","unique":true,"inverse":true,"annotations":{"EntGQL":{"Mapping":["githubAccount"]}}},{"name":"github_organization","type":"GithubOrganization","ref_name":"repositories","unique":true,"inverse":true,"annotations":{"EntGQL":{"Mapping":["githubOrganization"]}}}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"DESCRIPTION"}}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"User","config":{"Table":""},"edges":[{"name":"discord_accounts","type":"DiscordAccount","annotations":{"EntGQL":{"Mapping":["discordAccounts"]}}},{"name":"github_accounts","type":"GithubAccount","annotations":{"EntGQL":{"Mapping":["githubAccounts"]}}},{"name":"project_contributions","type":"ProjectContributor","annotations":{"EntGQL":{"Mapping":["projectContributions"]}}}],"fields":[{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"unique":true,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"USERNAME"}}},{"name":"avatar_url","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]}],"Features":["privacy","schema/snapshot"]}`

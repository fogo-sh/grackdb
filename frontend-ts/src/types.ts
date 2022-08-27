import { z } from "zod";
import {
  AuthProviderType,
  ProjectAssociationType,
  ProjectContributorRole,
  ProjectTechnologyAssociationType,
} from "./generated/graphql";

export const AuthProviderSchema = z.object({
  type: z.nativeEnum(AuthProviderType),
  url: z.string(),
});

export type AuthProvider = z.infer<typeof AuthProviderSchema>;

export const RepositorySchema = z.object({
  id: z.string(),
  name: z.string(),
  description: z.string().nullish(),
});

export type Repository = z.infer<typeof RepositorySchema>;

export const SiteSchema = z.object({
  id: z.string(),
  url: z.string(),
});

export type Site = z.infer<typeof SiteSchema>;

export const DiscordAccountSchema = z.object({
  id: z.string().nullish(),
  username: z.string(),
  discriminator: z.string(),
});

export type DiscordAccount = z.infer<typeof DiscordAccountSchema>;

export const DiscordBotSchema = z.object({
  id: z.string(),
  account: DiscordAccountSchema,
});

export type DiscordBot = z.infer<typeof DiscordBotSchema>;

export const TechnologySchema = z.object({
  id: z.string(),
  name: z.string(),
  colour: z.string().nullable(),
});

export type Technology = z.infer<typeof TechnologySchema>;

export const ProjectTechnologySchema = z.object({
  id: z.string().optional(),
  type: z.nativeEnum(ProjectTechnologyAssociationType).nullish(),
  technology: TechnologySchema,
});

export type ProjectTechnology = z.infer<typeof ProjectTechnologySchema>;

export const ProjectSchema = z.object({
  id: z.string(),
  name: z.string(),
  description: z.string().nullish(),
  startDate: z.string().nullish(),
  endDate: z.string().nullish(),
  technologies: z.array(ProjectTechnologySchema).optional(),
  repositories: z.array(RepositorySchema).optional(),
  sites: z.array(SiteSchema).optional(),
  discordBots: z.array(DiscordBotSchema).optional(),
});

export type Project = z.infer<typeof ProjectSchema>;

export const ProjectAssociationSchema = z.object({
  id: z.string(),
  type: z.nativeEnum(ProjectAssociationType),
  parent: ProjectSchema.nullish(),
  child: ProjectSchema.nullish(),
});

export type ProjectAssociation = z.infer<typeof ProjectAssociationSchema>;

export const ProjectContributionSchema = z.object({
  id: z.string().optional(),
  role: z.nativeEnum(ProjectContributorRole),
  project: ProjectSchema.nullish(),
});

export type ProjectContribution = z.infer<typeof ProjectContributionSchema>;

export const GithubAccountSchema = z.object({
  id: z.string(),
  username: z.string(),
});

export type GithubAccount = z.infer<typeof GithubAccountSchema>;

export const UserSchema = z.object({
  id: z.string().optional(),
  username: z.string(),
  avatarUrl: z.string().nullish(),
  githubAccounts: z.array(GithubAccountSchema).optional(),
  discordAccounts: z.array(DiscordAccountSchema).optional(),
  projectContributions: z.array(ProjectContributionSchema).optional(),
});

export type User = z.infer<typeof UserSchema>;

export const ProjectContributorSchema = ProjectContributionSchema.merge(
  z.object({
    user: UserSchema,
  })
);

export type ProjectContributor = z.infer<typeof ProjectContributorSchema>;

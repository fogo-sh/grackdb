import { z } from "zod";

export const TechnologySchema = z.object({
  id: z.string(),
  name: z.string(),
  colour: z.string().nullable(),
});

export type Technology = z.infer<typeof TechnologySchema>;

export const ProjectTechnologySchema = z.object({
  id: z.string().optional(),
  type: z.string().optional(), // TODO fetch enum from generated.ts
  technology: TechnologySchema,
});

export type ProjectTechnology = z.infer<typeof ProjectTechnologySchema>;

export const ProjectSchema = z.object({
  id: z.string(),
  name: z.string(),
  startDate: z.string().nullish(),
  endDate: z.string().nullish(),
  technologies: z.array(ProjectTechnologySchema).optional(),
});

export type Project = z.infer<typeof ProjectSchema>;

export const ProjectContributionSchema = z.object({
  id: z.string().optional(),
  role: z.string(), // TODO
  project: ProjectSchema,
});

export type ProjectContribution = z.infer<typeof ProjectContributionSchema>;

export const GithubAccountSchema = z.object({
  id: z.string(),
  username: z.string(),
});

export type GithubAccount = z.infer<typeof GithubAccountSchema>;

export const DiscordAccountSchema = z.object({
  id: z.string(),
  username: z.string(),
  discriminator: z.string(),
});

export type DiscordAccount = z.infer<typeof DiscordAccountSchema>;

export const UserSchema = z.object({
  id: z.string(),
  username: z.string(),
  avatarUrl: z.string().nullable(),
  githubAccounts: z.array(GithubAccountSchema).optional(),
  discordAccounts: z.array(DiscordAccountSchema).optional(),
  projectContributions: z.array(ProjectContributionSchema).optional(),
});

export type User = z.infer<typeof UserSchema>;

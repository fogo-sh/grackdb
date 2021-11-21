import React from "react";
import { useQuery } from "urql";
import { useParams } from "react-router-dom";

import { GithubAccountReference } from "../components/GithubAccount";
import { DiscordAccountReference } from "../components/DiscordAccount";
import { ProjectReference } from "../components/Project";
import { TechnologiesReference } from "../components/Technology";
import { enumValueToDisplayName } from "../utils";

const USERS_BY_USERNAME_QUERY = `
query UsersByUsername($username: String!) {
	users(where: {username: $username}) {
		edges {
			node {
				id
				username
				avatarUrl
				githubAccounts {
					id
					username
				}
				discordAccounts {
					id
					username
					discriminator
				}
				projectContributions {
					role
					project {
						id
						name
						technologies {
							technology {
								id
								name
								colour
							}
						}
					}
				}
			}
		}
	}
}
`;

export function UserPage() {
	const params = useParams();

	const [{ fetching, data }] = useQuery({
		query: USERS_BY_USERNAME_QUERY,
		variables: { username: params.username },
	});

	if (fetching) {
		return null;
	}

	const user = data.users.edges[0].node;

	return (
		<>
			<h1 className="text-center">{user.username}</h1>
			<h2>Accounts</h2>
			<div className="mx-2">
				{user.githubAccounts.map((githubAccount) => (
					<GithubAccountReference
						key={githubAccount.id}
						githubAccount={githubAccount}
						hasLink
					/>
				))}
				{user.discordAccounts.map((discordAccount) => (
					<DiscordAccountReference
						key={discordAccount.id}
						discordAccount={discordAccount}
					/>
				))}
			</div>
			<h2>Project Contributions</h2>
			<div className="mx-2">
				{user.projectContributions.map(({ role, project }) => (
					<ProjectReference key={project.id} project={project} hasLink>
						{({ projectName }) => (
							<>
								<div>
									{projectName} <i>({enumValueToDisplayName(role)})</i>
								</div>
								<TechnologiesReference technologies={project.technologies} />
							</>
						)}
					</ProjectReference>
				))}
			</div>
		</>
	);
}

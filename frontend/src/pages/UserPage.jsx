import React from "react";
import { useMutation, useQuery } from "urql";
import { useNavigate, useParams } from "react-router-dom";

import { GithubAccountReference } from "../components/GithubAccount";
import { DiscordAccountReference } from "../components/DiscordAccount";
import { ProjectReference } from "../components/Project";
import { TechnologiesReference } from "../components/Technology";
import { enumValueToDisplayName } from "../utils";
import { useAuth } from "../providers/AuthProvider";
import { useCallback } from "react";
import toast from "react-hot-toast";

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

const DELETE_USER_MUTATION = `
mutation DeleteUser($id: ID!) {
  deleteUser(id: $id) {
    id
  }
}
`;

export function UserPage() {
	const params = useParams();
	const navigate = useNavigate();

	const [, deleteUser] = useMutation(DELETE_USER_MUTATION);

	const { currentUser } = useAuth();

	const [{ fetching, data }] = useQuery({
		query: USERS_BY_USERNAME_QUERY,
		variables: { username: params.username },
	});

	const user = data?.users.edges[0].node;

	const onDeleteUser = useCallback(() => {
		deleteUser({ id: user.id });
		navigate("/users");
		toast.success(`Deleted ${user.username}`);
	});

	if (fetching) {
		return null;
	}

	return (
		<>
			<h1 className="text-center">{user.username}</h1>
			<h2>Accounts</h2>
			<div className="mx-2">
				{user.githubAccounts.length === 0 &&
					user.discordAccounts.length === 0 && <i>None</i>}
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
				{user.projectContributions.length === 0 && <i>None</i>}
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
			{currentUser && (
				<>
					<h2>Actions</h2>
					<button className="btn btn-primary" onClick={() => onDeleteUser()}>
						Delete User
					</button>
				</>
			)}
		</>
	);
}

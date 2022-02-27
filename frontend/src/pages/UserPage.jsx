import React, { useState } from "react";
import { useMutation, useQuery } from "urql";
import { useNavigate, useParams } from "react-router-dom";

import {
	CreateGithubAccountModal,
	GithubAccountReference,
} from "../components/GithubAccount";
import {
	CreateDiscordAccountModal,
	DiscordAccountReference,
} from "../components/DiscordAccount";
import { ProjectReference } from "../components/Project";
import { TechnologiesReference } from "../components/Technology";
import { enumValueToDisplayName } from "../utils";
import { useAuth } from "../providers/AuthProvider";
import { useCallback } from "react";
import toast from "react-hot-toast";
import { useErrorNotify } from "../hooks/useErrorNotify";
import { DeleteButton } from "../components/DeleteButton";

const USERS_BY_USERNAME_QUERY = /* GraphQL */ `
	query UsersByUsername($username: String!) {
		users(where: { username: $username }) {
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

const DELETE_USER_MUTATION = /* GraphQL */ `
	mutation DeleteUser($id: ID!) {
		deleteUser(id: $id) {
			id
		}
	}
`;

const DELETE_DISCORD_ACCOUNT_MUTATION = /* GraphQL */ `
	mutation DeleteDiscordAccount($id: ID!) {
		deleteDiscordAccount(id: $id) {
			id
		}
	}
`;

const DELETE_GITHUB_ACCOUNT_MUTATION = /* GraphQL */ `
	mutation DeleteGithubAccount($id: ID!) {
		deleteGithubAccount(id: $id) {
			id
		}
	}
`;

export function UserPage() {
	const params = useParams();
	const navigate = useNavigate();

	const [createDiscordAccountDialogOpen, setCreateDiscordAccountDialogOpen] =
		useState(false);
	const [createGithubAccountDialogOpen, setCreateGithubAccountDialogOpen] =
		useState(false);

	const [{ error: deleteUserError }, deleteUser] =
		useMutation(DELETE_USER_MUTATION);
	useErrorNotify(deleteUserError);

	const [{ error: deleteDiscordAccountError }, deleteDiscordAccount] =
		useMutation(DELETE_DISCORD_ACCOUNT_MUTATION);
	useErrorNotify(deleteDiscordAccountError);

	const [{ error: deleteGithubAccountError }, deleteGithubAccount] =
		useMutation(DELETE_GITHUB_ACCOUNT_MUTATION);
	useErrorNotify(deleteGithubAccountError);

	const { currentUser } = useAuth();

	const [{ fetching, data }, reexecuteUserQuery] = useQuery({
		query: USERS_BY_USERNAME_QUERY,
		variables: { username: params.username },
	});

	const refetchUser = () =>
		reexecuteUserQuery({ requestPolicy: "network-only" });

	const user = data?.users.edges[0].node;

	const onDeleteUser = useCallback(() => {
		if (confirm(`Are you sure you want to delete ${user.username}?`)) {
			deleteUser({ id: user.id });
			navigate("/users");
			toast.success(`Deleted ${user.username}`);
		}
	}, [user]);

	const onDeleteDiscordAccount = useCallback(
		(discordAccount) => {
			deleteDiscordAccount({ id: discordAccount.id });
			toast.success(`Deleted Discord Account from ${user.username}`);
		},
		[user]
	);

	const onDeleteGithubAccount = useCallback(
		(githubAccount) => {
			deleteGithubAccount({ id: githubAccount.id });
			toast.success(`Deleted GitHub Account from ${user.username}`);
		},
		[user]
	);

	if (fetching) {
		return null;
	}

	return (
		<>
			{user.avatarUrl !== null && (
				<img
					className="w-[10rem] h-[10rem] mx-auto border border-1 mt-5"
					src={user.avatarUrl}
					alt={user.username}
				/>
			)}
			<h1 className="text-center">{user.username}</h1>
			<h2>Accounts</h2>
			<div className="mx-2">
				{user.githubAccounts.length === 0 &&
					user.discordAccounts.length === 0 && <i>None</i>}
				{user.githubAccounts.map((githubAccount) => (
					<div key={githubAccount.id} className="flex gap-3">
						{currentUser && (
							<DeleteButton
								onClick={() => onDeleteGithubAccount(githubAccount)}
							/>
						)}
						<GithubAccountReference
							key={githubAccount.id}
							githubAccount={githubAccount}
							hasLink
						/>
					</div>
				))}
				{user.discordAccounts.map((discordAccount) => (
					<div key={discordAccount.id} className="flex gap-3">
						{currentUser && (
							<DeleteButton
								onClick={() => onDeleteDiscordAccount(discordAccount)}
							/>
						)}
						<DiscordAccountReference discordAccount={discordAccount} />
					</div>
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
					<div className="flex gap-2">
						<button className="btn btn-primary" onClick={() => onDeleteUser()}>
							Delete User
						</button>
						<button
							className="btn h-1/2"
							onClick={() => setCreateDiscordAccountDialogOpen(true)}
						>
							Associate Discord Account
						</button>
						<CreateDiscordAccountModal
							user={user}
							refetchUser={refetchUser}
							dialogOpen={createDiscordAccountDialogOpen}
							setDialogOpen={setCreateDiscordAccountDialogOpen}
						/>
						<button
							className="btn h-1/2"
							onClick={() => setCreateGithubAccountDialogOpen(true)}
						>
							Associate GitHub Account
						</button>
						<CreateGithubAccountModal
							user={user}
							refetchUser={refetchUser}
							dialogOpen={createGithubAccountDialogOpen}
							setDialogOpen={setCreateGithubAccountDialogOpen}
						/>
					</div>
				</>
			)}
		</>
	);
}

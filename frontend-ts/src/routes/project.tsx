import React from "react";
import ReactMarkdown from "react-markdown";
import { useMutation, useQuery } from "urql";
import { useNavigate, useParams } from "react-router-dom";

import {
	CreateProjectTechnologyModal,
	TechnologiesReference,
	TechnologyReference,
} from "../components/Technology";
import { enumValueToDisplayName } from "../utils";
import { ProjectDates, ProjectReference } from "../components/Project";
import { GithubRepositoryReference } from "../components/Repositories";
import { DiscordAccountReference } from "../components/DiscordAccount";
import { SiteReference } from "../components/Site";
import { UserReference } from "../components/User";
import { useAuth } from "../providers/AuthProvider";
import { useErrorNotify } from "../hooks/useErrorNotify";
import toast from "react-hot-toast";
import { useCallback } from "react";
import { useState } from "react";

const PROJECTS_BY_PROJECT_ID_QUERY = /* GraphQL */ `
	query ProjectsByProjectId($projectId: ID!) {
		projects(where: { id: $projectId }) {
			edges {
				node {
					id
					name
					description
					startDate
					endDate
					repositories {
						id
						name
						githubAccount {
							username
						}
						githubOrganization {
							name
						}
					}
					sites {
						id
						url
					}
					discordBots {
						id
						account {
							username
							discriminator
						}
					}
					parentProjects {
						id
						type
						parent {
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
					childProjects {
						id
						type
						child {
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
					contributors {
						id
						role
						user {
							username
						}
					}
					technologies {
						id
						type
						technology {
							id
							type
							name
							colour
						}
					}
				}
			}
		}
	}
`;

const DELETE_PROJECT_MUTATION = /* GraphQL */ `
	mutation DeleteProject($id: ID!) {
		deleteProject(id: $id) {
			id
		}
	}
`;

export function ProjectPage() {
	const params = useParams();
	const navigate = useNavigate();

	const [
		createProjectTechnologyDialogOpen,
		setCreateProjectTechnologyDialogOpen,
	] = useState(false);

	const [{ fetching, data }, reexecuteProjectQuery] = useQuery({
		query: PROJECTS_BY_PROJECT_ID_QUERY,
		variables: { projectId: params.projectId },
	});

	const refetchProject = () =>
		reexecuteProjectQuery({ requestPolicy: "network-only" });

	const [{ error: deleteProjectError }, deleteProject] = useMutation(
		DELETE_PROJECT_MUTATION
	);
	useErrorNotify(deleteProjectError);

	const { currentUser } = useAuth();

	const project = data?.projects.edges[0].node;

	const onDeleteProject = useCallback(() => {
		if (confirm(`Are you sure you want to delete ${project.name}?`)) {
			deleteProject({ id: project.id });
			navigate("/projects");
			toast.success(`Deleted ${project.name}`);
		}
	}, [project]);

	if (fetching) {
		return null;
	}

	return (
		<>
			<h1 className="text-center">{project.name}</h1>
			<p className="text-center mb-3">
				<ProjectDates project={project} />
			</p>
			<div className="flex h-full justify-around">
				{project.technologies.map(({ id, type, technology }) => (
					<div key={id} className="flex">
						<TechnologyReference key={technology.id} technology={technology}>
							{({ circle, name }) => (
								<>
									{circle}{" "}
									<span>
										{enumValueToDisplayName(type)} {name}
									</span>
								</>
							)}
						</TechnologyReference>
					</div>
				))}
			</div>

			<div className="my-1">
				{project.repositories.map((repository) => (
					<GithubRepositoryReference
						key={repository.id}
						repository={repository}
						hasLink
					/>
				))}
				{project.sites.map((site) => (
					<SiteReference key={site.id} site={site} hasLink />
				))}
				{project.discordBots.map((discordBot) => (
					<DiscordAccountReference
						key={discordBot.id}
						discordAccount={discordBot.account}
						hasLink
					/>
				))}
			</div>

			<ReactMarkdown className="mt-3 mb-4">{project.description}</ReactMarkdown>

			<h2>Contributors</h2>
			<div className="mx-2">
				{(!project.contributors || project.contributors.length === 0) && (
					<i>None</i>
				)}
				{project.contributors?.map((contributor) => (
					<UserReference key={contributor.id} user={contributor.user} hasLink>
						{({ userName }) => (
							<>
								{userName} <i>({enumValueToDisplayName(contributor.role)})</i>
							</>
						)}
					</UserReference>
				))}
			</div>

			{project.parentProjects.length !== 0 && (
				<>
					<h2>Parent Projects</h2>
					<div className="mx-2">
						{project.parentProjects.map((parentProject) => (
							<ProjectReference
								key={parentProject.id}
								project={parentProject.parent}
								hasLink
							>
								{({ projectName }) => (
									<>
										<span>
											{projectName}{" "}
											<i>
												({project.name}{" "}
												{enumValueToDisplayName(parentProject.type)})
											</i>
										</span>
										<TechnologiesReference
											technologies={parentProject.parent.technologies}
										/>
									</>
								)}
							</ProjectReference>
						))}
					</div>
				</>
			)}

			{project.childProjects.length !== 0 && (
				<>
					<h2>Child Projects</h2>
					<div className="mx-2">
						{project.childProjects.map((childProject) => (
							<ProjectReference
								key={childProject.id}
								project={childProject.child}
								hasLink
							>
								{({ projectName }) => (
									<>
										<span>
											{projectName}{" "}
											<i>
												({enumValueToDisplayName(childProject.type)}{" "}
												{project.name})
											</i>
										</span>
										<TechnologiesReference
											technologies={childProject.child.technologies}
										/>
									</>
								)}
							</ProjectReference>
						))}
					</div>
				</>
			)}
			{currentUser && (
				<>
					<h2>Actions</h2>
					<div className="flex gap-2">
						<button
							className="btn btn-primary"
							onClick={() => onDeleteProject()}
						>
							Delete Project
						</button>
						<button
							className="btn btn-primary"
							onClick={() => setCreateProjectTechnologyDialogOpen(true)}
						>
							Associate Technology
						</button>
						<CreateProjectTechnologyModal
							project={project}
							refetchProject={refetchProject}
							dialogOpen={createProjectTechnologyDialogOpen}
							setDialogOpen={setCreateProjectTechnologyDialogOpen}
						/>
					</div>
				</>
			)}
		</>
	);
}

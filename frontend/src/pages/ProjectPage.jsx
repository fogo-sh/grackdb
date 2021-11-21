import React from "react";
import { useQuery } from "urql";
import { useParams } from "react-router-dom";

import {
	TechnologiesReference,
	TechnologyReference,
} from "../components/Technology";
import { enumValueToDisplayName } from "../utils";
import { ProjectDates } from "../components/Project";

const PROJECTS_BY_PROJECT_ID_QUERY = `
query ProjectsByProjectId($projectId: ID!) {
	projects(where: { id: $projectId }) {
		edges {
			node {
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
					}
				}
				parentProjects {
					type
					parent {
						name
					}
				}
				childProjects {
					type
					child {
						name
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

export function ProjectPage() {
	const params = useParams();

	const [{ fetching, data }] = useQuery({
		query: PROJECTS_BY_PROJECT_ID_QUERY,
		variables: { projectId: params.projectId },
	});

	if (fetching) {
		return null;
	}

	const project = data.projects.edges[0].node;

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
		</>
	);
}

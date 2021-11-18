import React from "react";
import { useQuery } from "urql";
import { useParams } from "react-router-dom";

import { ProjectReference } from "../components/Project";
import { TechnologiesReference } from "../components/Technology";

const PROJECTS_BY_PROJECT_ID_QUERY = `
query ProjectsByProjectId($projectId: ID!) {
	projects(where: {id: $projectId}) {
		edges {
			node {
				name
				technologies {
					technology {
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
		<ProjectReference project={project}>
			<TechnologiesReference technologies={project.technologies} />
		</ProjectReference>
	);
}

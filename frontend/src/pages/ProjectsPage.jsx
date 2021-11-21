import React from "react";
import { useQuery } from "urql";

import { ProjectReference } from "../components/Project";
import { TechnologiesReference } from "../components/Technology";

const PROJECTS_QUERY = `
query Projects {
	projects {
		edges {
			node {
				id
				name
				technologies {
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

export function ProjectsPage() {
	const [{ fetching, data }] = useQuery({
		query: PROJECTS_QUERY,
	});

	if (fetching) {
		return null;
	}

	const projects = data.projects.edges.map((edge) => edge.node);

	return (
		<>
			<h2>Projects</h2>
			<div className="mx-2">
				{projects.map((project) => (
					<ProjectReference key={project.id} project={project} hasLink>
						{({ projectName }) => (
							<>
								{projectName}
								<TechnologiesReference technologies={project.technologies} />
							</>
						)}
					</ProjectReference>
				))}
			</div>
		</>
	);
}

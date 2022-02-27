import React from "react";
import { useState } from "react";
import { useQuery } from "urql";

import { useAuth } from "../providers/AuthProvider";
import { CreateProjectModal, ProjectReference } from "../components/Project";
import { TechnologiesReference } from "../components/Technology";

const PROJECTS_QUERY = /* GraphQL */ `
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
	const [dialogOpen, setDialogOpen] = useState(false);

	const { currentUser } = useAuth();

	const [{ fetching, data }] = useQuery({
		query: PROJECTS_QUERY,
	});

	if (fetching) {
		return null;
	}

	const projects = data.projects.edges.map((edge) => edge.node);

	return (
		<>
			<div className="flex justify-between items-center">
				<h2 className="h-full my-0">Projects</h2>
				{currentUser && (
					<>
						<button className="btn h-1/2" onClick={() => setDialogOpen(true)}>
							Create
						</button>
						<CreateProjectModal
							dialogOpen={dialogOpen}
							setDialogOpen={setDialogOpen}
						/>
					</>
				)}
			</div>
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

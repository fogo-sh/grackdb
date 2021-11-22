import React from "react";
import { useQuery } from "urql";

import { ProjectReference } from "../components/Project";
import { TechnologiesReference } from "../components/Technology";
import { UserReference } from "../components/User";

const HOMEPAGE_QUERY = `
query Homepage {
	users {
		edges {
			node {
				id
				username
				avatarUrl
			}
		}
	}
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

export function HomepagePage() {
	const [{ fetching, data }] = useQuery({
		query: HOMEPAGE_QUERY,
	});

	if (fetching) {
		return null;
	}

	const users = data.users.edges.map((edge) => edge.node);
	const projects = data.projects.edges.map((edge) => edge.node);

	return (
		<>
			<h1 className="text-center">GrackDB</h1>
			<p className="text-center">a database</p>
			<h2>Users</h2>
			<div className="mx-2">
				{users.length === 0 && <i>None</i>}
				{users.map((user) => (
					<UserReference key={user.id} user={user} hasLink>
						{({ userName }) => userName}
					</UserReference>
				))}
			</div>
			<h2>Projects</h2>
			<div className="mx-2">
				{projects.length === 0 && <i>None</i>}
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

import React from "react";
import { Link } from "react-router-dom";

export function ProjectReference({ project, children, hasLink = false }) {
	return (
		<div className="flex justify-between my-2 h-6 items-center">
			{hasLink ? (
				<Link to={`/project/${project.id}`}>{project.name}</Link>
			) : (
				project.name
			)}
			{children}
		</div>
	);
}

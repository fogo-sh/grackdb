import React from "react";
import { format } from "date-fns";
import { Link } from "react-router-dom";

export function ProjectReference({ project, hasLink = false, children }) {
	const projectName = hasLink ? (
		<Link to={`/project/${project.id}`}>{project.name}</Link>
	) : (
		project.name
	);

	return (
		<div className="flex justify-between my-2 h-6 items-center">
			{children({ projectName })}
		</div>
	);
}

const DATE_FORMAT = "LLLL do, Y";

export function ProjectDates({ project }) {
	const startDate = project.startDate ? format(new Date(project.startDate), DATE_FORMAT) : '???';
	const endDate = project.endDate ? format(new Date(project.endDate), DATE_FORMAT) : '???';

	return <>{startDate} - {endDate}</>;
}

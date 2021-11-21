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
	const hasStartDate = project.startDate !== null;
	const hasEndDate = project.endDate !== null;

	const startDate = hasStartDate
		? format(new Date(project.startDate), DATE_FORMAT)
		: "Present";
	const endDate = hasEndDate
		? format(new Date(project.endDate), DATE_FORMAT)
		: "Present";

	return (
		<>
			{startDate} - {endDate}
		</>
	);
}

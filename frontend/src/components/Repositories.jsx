import React from "react";
import { FaExternalLinkAlt, FaGithub } from "react-icons/fa";

export function GithubRepositoryReference({ repository, hasLink = false }) {
	const parent =
		repository.githubAccount?.username ?? repository.githubOrganization?.name;

	const path = `${parent}/${repository.name}`;

	return (
		<div className="flex my-2 items-center">
			<FaGithub className="mr-1" />
			{hasLink ? (
				<a href={`https://github.com/${path}`} className="flex items-center">
					{path} <FaExternalLinkAlt className="ml-1 text-xs" />
				</a>
			) : (
				path
			)}
		</div>
	);
}

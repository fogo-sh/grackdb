import React from "react";
import { FaExternalLinkAlt, FaGithub } from "react-icons/fa";

export function GithubAccountReference({ githubAccount, hasLink = false }) {
	return (
		<div className="flex my-2 items-center">
			<FaGithub className="mr-1" />
			{hasLink ? (
				<>
					<a href={`https://github.com/${githubAccount.username}`} className="flex items-center">
						{githubAccount.username}
						<FaExternalLinkAlt className="ml-1 text-xs"/>
					</a>
				</>
			) : (
				githubAccount.username
			)}
		</div>
	);
}

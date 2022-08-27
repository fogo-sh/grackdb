import { FaExternalLinkAlt, FaGithub } from "react-icons/fa";
import { Repository } from "~/types";

export function GithubRepositoryReference({
  repository,
  hasLink = false,
}: {
  repository: Repository;
  hasLink: boolean;
}) {
  const parent =
    repository.githubAccount?.username ??
    repository.githubOrganization?.name ??
    "???";

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

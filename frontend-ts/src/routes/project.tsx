import { useLoaderData } from "react-router-dom";
import ReactMarkdown from "react-markdown";
import { enumValueToDisplayName } from "~/utils";
import {
  TechnologiesReference,
  TechnologyReference,
} from "~/components/Technology";
import { ProjectDates, ProjectReference } from "~/components/Project";
import { GithubRepositoryReference } from "~/components/Repositories";
import { DiscordAccountReference } from "~/components/DiscordAccount";
import { SiteReference } from "~/components/Site";
import { UserReference } from "~/components/User";

export function ProjectPage() {
  const { project } = useLoaderData();

  const currentUser = undefined; // TODO

  return (
    <>
      <h1 className="text-center">{project.name}</h1>
      <p className="text-center mb-3">
        <ProjectDates project={project} />
      </p>
      <div className="flex h-full justify-around">
        {project.technologies.map(({ id, type, technology }) => (
          <div key={id} className="flex">
            <TechnologyReference key={technology.id} technology={technology}>
              {({ circle, name }) => (
                <>
                  {circle}{" "}
                  <span>
                    {enumValueToDisplayName(type)} {name}
                  </span>
                </>
              )}
            </TechnologyReference>
          </div>
        ))}
      </div>

      <div className="my-1">
        {project.repositories.map((repository) => (
          <GithubRepositoryReference
            key={repository.id}
            repository={repository}
            hasLink
          />
        ))}
        {project.sites.map((site) => (
          <SiteReference key={site.id} site={site} hasLink />
        ))}
        {project.discordBots.map((discordBot) => (
          <DiscordAccountReference
            key={discordBot.id}
            discordAccount={discordBot.account}
            hasLink
          />
        ))}
      </div>

      <ReactMarkdown className="mt-3 mb-4">{project.description}</ReactMarkdown>

      <h2>Contributors</h2>
      <div className="mx-2">
        {(!project.contributors || project.contributors.length === 0) && (
          <i>None</i>
        )}
        {project.contributors?.map((contributor) => (
          <UserReference key={contributor.id} user={contributor.user} hasLink>
            {({ userName }) => (
              <>
                {userName} <i>({enumValueToDisplayName(contributor.role)})</i>
              </>
            )}
          </UserReference>
        ))}
      </div>

      {project.parentProjects.length !== 0 && (
        <>
          <h2>Parent Projects</h2>
          <div className="mx-2">
            {project.parentProjects.map((parentProject) => (
              <ProjectReference
                key={parentProject.id}
                project={parentProject.parent}
                hasLink
              >
                {({ projectName }) => (
                  <>
                    <span>
                      {projectName}{" "}
                      <i>
                        ({project.name}{" "}
                        {enumValueToDisplayName(parentProject.type)})
                      </i>
                    </span>
                    <TechnologiesReference
                      technologies={parentProject.parent.technologies}
                    />
                  </>
                )}
              </ProjectReference>
            ))}
          </div>
        </>
      )}

      {project.childProjects.length !== 0 && (
        <>
          <h2>Child Projects</h2>
          <div className="mx-2">
            {project.childProjects.map((childProject) => (
              <ProjectReference
                key={childProject.id}
                project={childProject.child}
                hasLink
              >
                {({ projectName }) => (
                  <>
                    <span>
                      {projectName}{" "}
                      <i>
                        ({enumValueToDisplayName(childProject.type)}{" "}
                        {project.name})
                      </i>
                    </span>
                    <TechnologiesReference
                      technologies={childProject.child.technologies}
                    />
                  </>
                )}
              </ProjectReference>
            ))}
          </div>
        </>
      )}
      {currentUser && (
        <>
          <h2>Actions</h2>
          <div className="flex gap-2">
            <button className="btn btn-primary" onClick={() => alert(true)}>
              Delete Project
            </button>
            <button className="btn btn-primary" onClick={() => alert(true)}>
              Associate Technology
            </button>
          </div>
        </>
      )}
    </>
  );
}

import { LoaderFunction, useLoaderData } from "react-router-dom";
import { z } from "zod";
import { ProjectReference } from "~/components/Project";
import { TechnologiesReference } from "~/components/Technology";
import { useProjectsQuery } from "~/generated/graphql";
import { dataSource, queryClient } from "~/query";
import { ProjectSchema } from "~/types";

const LoaderDataSchema = z.object({
  projects: z.array(ProjectSchema),
});

type LoaderData = z.infer<typeof LoaderDataSchema>;

export const loader: LoaderFunction = async () => {
  const projectsQuery = await queryClient.fetchQuery(
    useProjectsQuery.getKey(),
    async () => await useProjectsQuery.fetcher(dataSource)()
  );

  const projects = projectsQuery.projects?.edges?.map((edge) => edge?.node);

  const data: LoaderData = LoaderDataSchema.parse({ projects });
  return data;
};

export function ProjectsPage() {
  const { projects } = useLoaderData() as LoaderData;

  const currentUser = undefined; // TODO

  return (
    <>
      <div className="flex justify-between items-center">
        <h2 className="h-full my-0">Projects</h2>
        {currentUser && (
          <>
            <button className="btn h-1/2" onClick={() => alert(true)}>
              Create
            </button>
          </>
        )}
      </div>
      <div className="mx-2">
        {projects.map((project) => (
          <ProjectReference key={project.id} project={project} hasLink>
            {({ projectName }) => (
              <>
                {projectName}
                <TechnologiesReference
                  technologies={project.technologies ?? []}
                />
              </>
            )}
          </ProjectReference>
        ))}
      </div>
    </>
  );
}

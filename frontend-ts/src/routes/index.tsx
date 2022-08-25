import { LoaderFunction, useLoaderData } from "react-router-dom";
import { z } from "zod";
import { dataSource, queryClient } from "~/query";
import { ProjectSchema, UserSchema } from "~/types";
import { useHomepageQuery } from "~/generated/graphql";
import { UserReference } from "~/components/User";
import { ProjectReference } from "~/components/Project";
import { TechnologiesReference } from "~/components/Technology";

const LoaderDataSchema = z.object({
  users: z.array(UserSchema),
  projects: z.array(ProjectSchema),
});

type LoaderData = z.infer<typeof LoaderDataSchema>;

export const loader: LoaderFunction = async () => {
  const homepageQuery = await queryClient.fetchQuery(
    useHomepageQuery.getKey(),
    async () => await useHomepageQuery.fetcher(dataSource)()
  );

  const users = homepageQuery.users?.edges?.map((edge) => edge?.node);
  const projects = homepageQuery.projects?.edges?.map((edge) => edge?.node);

  const data: LoaderData = LoaderDataSchema.parse({ users, projects });
  return data;
};

export function IndexPage() {
  const { users, projects } = useLoaderData() as LoaderData;

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

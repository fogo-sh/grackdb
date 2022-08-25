import { LoaderFunction, useLoaderData } from "react-router-dom";

import { GithubAccountReference } from "~/components/GithubAccount";
import { DiscordAccountReference } from "~/components/DiscordAccount";
import { ProjectReference } from "~/components/Project";
import { TechnologiesReference } from "~/components/Technology";
import { enumValueToDisplayName } from "~/utils";
import { DeleteButton } from "~/components/DeleteButton";
import { UserSchema } from "~/types";
import { dataSource, queryClient } from "~/query";
import { z } from "zod";
import { useUsersByUsernameQuery } from "~/generated/graphql";
import invariant from "tiny-invariant";

const LoaderDataSchema = z.object({
  user: UserSchema,
});

type LoaderData = z.infer<typeof LoaderDataSchema>;

export const loader: LoaderFunction = async ({ params }) => {
  const username = params.username;
  invariant(username);

  const userQuery = await queryClient.fetchQuery(
    useUsersByUsernameQuery.getKey({ username }),
    async () =>
      await useUsersByUsernameQuery.fetcher(dataSource, { username })()
  );

  const user = userQuery.users?.edges?.map((edge) => edge?.node)[0];

  const data: LoaderData = LoaderDataSchema.parse({ user });
  return data;
};

export function UserPage() {
  const currentUser = undefined; // TODO
  const { user } = useLoaderData() as LoaderData;

  return (
    <>
      {user.avatarUrl !== null && (
        <img
          className="w-[10rem] h-[10rem] mx-auto border border-1 mt-5"
          src={user.avatarUrl}
          alt={user.username}
        />
      )}
      <h1 className="text-center">{user.username}</h1>
      <h2>Accounts</h2>
      <div className="mx-2">
        {user.githubAccounts?.length === 0 &&
          user.discordAccounts?.length === 0 && <i>None</i>}
        {user.githubAccounts?.map((githubAccount) => (
          <div key={githubAccount.id} className="flex gap-3">
            {currentUser && <DeleteButton onClick={() => alert(true)} />}
            <GithubAccountReference
              key={githubAccount.id}
              githubAccount={githubAccount}
              hasLink
            />
          </div>
        ))}
        {user.discordAccounts?.map((discordAccount) => (
          <div key={discordAccount.id} className="flex gap-3">
            {currentUser && <DeleteButton onClick={() => alert(true)} />}
            <DiscordAccountReference discordAccount={discordAccount} />
          </div>
        ))}
      </div>
      <h2>Project Contributions</h2>
      <div className="mx-2">
        {user.projectContributions?.length === 0 && <i>None</i>}
        {user.projectContributions?.map(({ role, project }) => (
          <ProjectReference key={project.id} project={project} hasLink>
            {({ projectName }) => (
              <>
                <div>
                  {projectName} <i>({enumValueToDisplayName(role)})</i>
                </div>
                <TechnologiesReference
                  technologies={project.technologies ?? []}
                />
              </>
            )}
          </ProjectReference>
        ))}
      </div>
      {currentUser && (
        <>
          <h2>Actions</h2>
          <div className="flex gap-2">
            <button className="btn btn-primary" onClick={() => alert(true)}>
              Delete User
            </button>
            <button className="btn h-1/2" onClick={() => alert(true)}>
              Associate Discord Account
            </button>
            <button className="btn h-1/2" onClick={() => alert(true)}>
              Associate GitHub Account
            </button>
          </div>
        </>
      )}
    </>
  );
}

import { Link, LoaderFunction, Outlet, useLoaderData } from "react-router-dom";
import { z } from "zod";
import invariant from "tiny-invariant";
import { enumValueToDisplayName } from "~/utils";
import { UserSchema } from "~/types";
import { dataSource, queryClient } from "~/query";
import { DeleteButton } from "~/components/DeleteButton";
import { useUsersByUsernameQuery } from "~/generated/graphql";
import { GithubAccountReference } from "~/components/GithubAccount";
import { DiscordAccountReference } from "~/components/DiscordAccount";
import { ProjectReference } from "~/components/Project";
import { TechnologiesReference } from "~/components/Technology";
import { useCurrentUser } from "./root";

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
  const { user } = useLoaderData() as LoaderData;
  const currentUser = useCurrentUser();

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
            {currentUser && (
              <Link to={`./delete/github-account/${githubAccount.id}`}>
                <DeleteButton />
              </Link>
            )}
            <GithubAccountReference
              key={githubAccount.id}
              githubAccount={githubAccount}
              hasLink
            />
          </div>
        ))}
        {user.discordAccounts?.map((discordAccount) => (
          <div key={discordAccount.id} className="flex gap-3">
            {currentUser && (
              <Link to={`./delete/discord-account/${discordAccount.id}`}>
                <DeleteButton />
              </Link>
            )}
            <DiscordAccountReference discordAccount={discordAccount} />
          </div>
        ))}
      </div>
      <h2>Project Contributions</h2>
      <div className="mx-2">
        {user.projectContributions?.length === 0 && <i>None</i>}
        {user.projectContributions?.map(({ role, project }) => (
          <ProjectReference key={project?.id} project={project} hasLink>
            {({ projectName }) => (
              <>
                <div>
                  {projectName} <i>({enumValueToDisplayName(role)})</i>
                </div>
                <TechnologiesReference
                  technologies={project?.technologies ?? []}
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
            <Link to="./delete">
              <button className="btn btn-primary">Delete User</button>
            </Link>
            <Link to="./associate/discord-account">
              <button className="btn btn-primary">
                Associate Discord Account
              </button>
            </Link>
            <Link to="./associate/github-account">
              <button className="btn btn-primary">
                Associate GitHub Account
              </button>
            </Link>
          </div>
        </>
      )}
      <Outlet />
    </>
  );
}

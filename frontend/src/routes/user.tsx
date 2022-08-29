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
import { FaEdit } from "react-icons/fa";
import { useState } from "react";

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

  const [editing, setEditing] = useState(false);

  return (
    <>
      {user.avatarURL !== null && (
        <img
          className="w-[10rem] h-[10rem] mx-auto border border-1 mt-5"
          src={user.avatarURL}
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
            {editing && (
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
            {editing && (
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
                <div className="flex items-center gap-x-3">
                  {editing && (
                    <Link to={`./delete/contribution/${project?.id}`}>
                      <DeleteButton />
                    </Link>
                  )}
                  <div>
                    {projectName} <i>({enumValueToDisplayName(role)})</i>
                  </div>
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
          <div className="flex flex-wrap gap-2">
            <button
              className="btn float-right flex items-center gap-x-2"
              onClick={() => setEditing(!editing)}
            >
              <FaEdit /> {editing ? "Stop Editing" : "Edit"}
            </button>
            {editing && (
              <>
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
                <Link to="./associate/project-contribution">
                  <button className="btn btn-primary">
                    Associate Project Contribution
                  </button>
                </Link>
              </>
            )}
          </div>
        </>
      )}
      <Outlet />
    </>
  );
}

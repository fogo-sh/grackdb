import { ActionFunction, Form, redirect, useNavigate } from "react-router-dom";
import invariant from "tiny-invariant";
import { Modal } from "~/components/Modal";
import {
  useDeleteDiscordAccountMutation,
  useDeleteGithubAccountMutation,
  useUserIdFromUsernameQuery,
} from "~/generated/graphql";
import { dataSource, queryClient } from "~/query";

export const action: ActionFunction = async ({ params }) => {
  const username = params.username;
  invariant(username);

  const userIdQuery = await queryClient.fetchQuery(
    useUserIdFromUsernameQuery.getKey({ username }),
    async () =>
      await useUserIdFromUsernameQuery.fetcher(dataSource, { username })()
  );

  const id = userIdQuery?.users?.edges?.[0]?.node?.id;
  invariant(id);

  const thing = params.thing;
  invariant(thing);

  const thingId = params.thingId;
  invariant(thingId);

  if (thing === "github-account") {
    await queryClient.fetchQuery(
      useDeleteGithubAccountMutation.getKey(),
      async () =>
        await useDeleteGithubAccountMutation.fetcher(dataSource, {
          id: thingId,
        })()
    );
    return redirect(`/user/${username}`);
  }

  if (thing === "discord-account") {
    await queryClient.fetchQuery(
      useDeleteDiscordAccountMutation.getKey(),
      async () =>
        await useDeleteDiscordAccountMutation.fetcher(dataSource, {
          id: thingId,
        })()
    );
    return redirect(`/user/${username}`);
  }

  throw new Error("Unknown thing");
};

export function UserDeleteThing() {
  const navigate = useNavigate();

  return (
    <Modal title="Are you sure?" close={() => navigate("..")}>
      <Form method="post" className="flex flex-col gap-4">
        <input className="btn" type="submit" value="Delete" />
      </Form>
    </Modal>
  );
}

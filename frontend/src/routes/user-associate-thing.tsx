import {
  ActionFunction,
  Form,
  redirect,
  useNavigate,
  useParams,
} from "react-router-dom";
import invariant from "tiny-invariant";
import { z } from "zod";
import { Input } from "~/components/Form";
import { Modal } from "~/components/Modal";
import {
  useCreateDiscordAccountMutation,
  useCreateGithubAccountMutation,
  useUserIdFromUsernameQuery,
} from "~/generated/graphql";
import { dataSource, queryClient } from "~/query";

const CreateGithubAccountSchema = z.object({
  username: z.string(),
  owner: z.number(),
});

const CreateDiscordAccountSchema = z.object({
  username: z.string(),
  discriminator: z.string(),
  discordId: z.string(),
  owner: z.number(),
});

export const action: ActionFunction = async ({ request, params }) => {
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

  const formData = await request.formData();
  const entries = Object.fromEntries(formData);

  if (thing === "github-account") {
    const data = CreateGithubAccountSchema.parse({
      ...entries,
      owner: Number(id),
    });
    await queryClient.fetchQuery(
      useCreateGithubAccountMutation.getKey(),
      async () =>
        await useCreateGithubAccountMutation.fetcher(dataSource, data)()
    );
    return redirect(`/user/${username}`);
  }

  if (thing === "discord-account") {
    const data = CreateDiscordAccountSchema.parse({
      ...entries,
      owner: Number(id),
    });
    await queryClient.fetchQuery(
      useCreateDiscordAccountMutation.getKey(),
      async () =>
        await useCreateDiscordAccountMutation.fetcher(dataSource, data)()
    );
    return redirect(`/user/${username}`);
  }

  throw new Error("Unknown thing");
};

function AssociateDiscord() {
  return (
    <>
      <Input id="username" display="Discord Username" />
      <Input id="discriminator" display="Discord Discriminator" />
      <Input id="discordId" display="Discord ID" />
    </>
  );
}

function AssociateGitHub() {
  return (
    <>
      <Input id="username" display="GitHub Username" />
    </>
  );
}

export function UserAssociateThing() {
  const navigate = useNavigate();
  const params = useParams();
  const thing = params.thing;
  const username = params.username;

  return (
    <Modal title="Associate" close={() => navigate(-1)}>
      <Form method="post" className="flex flex-col gap-4">
        {thing === "discord-account" && <AssociateDiscord />}
        {thing === "github-account" && <AssociateGitHub />}
        <input className="btn" type="submit" value="Associate" />
      </Form>
    </Modal>
  );
}

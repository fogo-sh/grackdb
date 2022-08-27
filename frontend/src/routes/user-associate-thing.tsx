import {
  ActionFunction,
  Form,
  LoaderFunction,
  redirect,
  useLoaderData,
  useNavigate,
  useParams,
} from "react-router-dom";
import invariant from "tiny-invariant";
import { z } from "zod";
import { Input, SelectInput } from "~/components/Form";
import { Modal } from "~/components/Modal";
import {
  ProjectContributorRole,
  useCreateDiscordAccountMutation,
  useCreateGithubAccountMutation,
  useCreateProjectContributorMutation,
  useProjectsQuery,
  useUserIdFromUsernameQuery,
} from "~/generated/graphql";
import { dataSource, queryClient } from "~/query";

const LoaderDataSchema = z.object({
  projectOptions: z.array(z.object({ id: z.string(), value: z.string() })),
});

type LoaderData = z.infer<typeof LoaderDataSchema>;

export const loader: LoaderFunction = async ({ params }) => {
  let projectOptions = [] as any;

  if (params.thing === "project-contribution") {
    const projectsQuery = await queryClient.fetchQuery(
      useProjectsQuery.getKey(),
      async () => await useProjectsQuery.fetcher(dataSource)()
    );

    const projects =
      projectsQuery.projects?.edges?.map((edge) => edge?.node) ?? [];

    projectOptions = projects.map((project) => ({
      id: project?.id,
      value: project?.name,
    }));
  }

  const data: LoaderData = LoaderDataSchema.parse({ projectOptions });

  return data;
};

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

const CreateProjectContributionSchema = z.object({
  role: z.nativeEnum(ProjectContributorRole),
  project: z.number(),
  user: z.number(),
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

  if (thing === "project-contribution") {
    const data = CreateProjectContributionSchema.parse({
      ...entries,
      user: Number(id),
      project: Number(entries.project),
    });
    await queryClient.fetchQuery(
      useCreateProjectContributorMutation.getKey(),
      async () =>
        await useCreateProjectContributorMutation.fetcher(dataSource, data)()
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

const projectContributorRoleOptions = Object.entries(
  ProjectContributorRole
).map(([k, v]) => ({ id: v, value: k }));

function AssociateProjectContribution() {
  const { projectOptions } = useLoaderData() as LoaderData;

  return (
    <>
      <SelectInput
        id="role"
        display="Role"
        options={projectContributorRoleOptions}
      />
      <SelectInput id="project" display="Project" options={projectOptions} />
    </>
  );
}

export function UserAssociateThing() {
  const navigate = useNavigate();
  const params = useParams();
  const thing = params.thing;

  return (
    <Modal title="Associate" close={() => navigate(-1)}>
      <Form method="post" className="flex flex-col gap-4">
        {thing === "discord-account" && <AssociateDiscord />}
        {thing === "github-account" && <AssociateGitHub />}
        {thing === "project-contribution" && <AssociateProjectContribution />}
        <input className="btn" type="submit" value="Associate" />
      </Form>
    </Modal>
  );
}

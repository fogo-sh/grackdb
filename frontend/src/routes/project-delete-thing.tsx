import {
  ActionFunction,
  Form,
  redirect,
  useNavigate,
  useParams,
} from "react-router-dom";
import invariant from "tiny-invariant";
import { Modal } from "~/components/Modal";
import {
  useDeleteProjectAssociationMutation,
  useDeleteProjectContributorMutation,
  useDeleteProjectTechnologyMutation,
  useRemoveRepositoryProjectMutation,
} from "~/generated/graphql";
import { dataSource, queryClient } from "~/query";

export const action: ActionFunction = async ({ params }) => {
  const id = params.id;
  invariant(id);

  const thing = params.thing;
  invariant(thing);

  const thingId = params.thingId;
  invariant(thingId);

  if (thing === "technology-reference") {
    await queryClient.fetchQuery(
      useDeleteProjectTechnologyMutation.getKey(),
      async () =>
        await useDeleteProjectTechnologyMutation.fetcher(dataSource, {
          id: thingId,
        })()
    );
    return redirect(`/project/${id}`);
  }

  if (thing === "repository-reference") {
    // TODO fix this, might require nint changes / context
    throw new Error("Project reference deletion not supported yet");
    /*
    await queryClient.fetchQuery(
      useRemoveRepositoryProjectMutation.getKey(),
      async () =>
        await useRemoveRepositoryProjectMutation.fetcher(dataSource, {
          id: thingId,
        })()
    );
    return redirect(`/project/${id}`);
    */
  }

  if (thing === "discord-bot-reference") {
    // TODO fix this, might require nint changes / context
    throw new Error("Discord bot reference deletion not supported yet");
  }

  if (thing === "contributor") {
    await queryClient.fetchQuery(
      useDeleteProjectContributorMutation.getKey(),
      async () =>
        await useDeleteProjectContributorMutation.fetcher(dataSource, {
          id: thingId,
        })()
    );
    return redirect(`/project/${id}`);
  }

  if (thing === "project-reference") {
    await queryClient.fetchQuery(
      useDeleteProjectAssociationMutation.getKey(),
      async () =>
        await useDeleteProjectAssociationMutation.fetcher(dataSource, {
          id: thingId,
        })()
    );
    return redirect(`/project/${id}`);
  }

  throw new Error("Unknown thing");
};

export function ProjectDeleteThing() {
  const navigate = useNavigate();
  const id = useParams();
  invariant(id);

  return (
    <Modal title="Are you sure?" close={() => navigate(-1)}>
      <Form method="post" className="flex flex-col gap-4">
        <input className="btn" type="submit" value="Delete" />
      </Form>
    </Modal>
  );
}

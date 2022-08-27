import {
  ActionFunction,
  Form,
  redirect,
  useNavigate,
  useParams,
} from "react-router-dom";
import invariant from "tiny-invariant";
import { Modal } from "~/components/Modal";
import { useDeleteProjectTechnologyMutation } from "~/generated/graphql";
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
    alert("todo");
    return;
  }

  if (thing === "discord-bot-reference") {
    alert("todo");
    return;
  }

  if (thing === "contributor") {
    alert("todo");
    return;
  }

  if (thing === "project-reference") {
    alert("todo");
    return;
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

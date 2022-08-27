import { ActionFunction, Form, redirect, useNavigate } from "react-router-dom";
import invariant from "tiny-invariant";
import { Modal } from "~/components/Modal";
import { useDeleteProjectMutation } from "~/generated/graphql";
import { dataSource, queryClient } from "~/query";

export const action: ActionFunction = async ({ params }) => {
  const id = params.id;
  invariant(id);

  await queryClient.fetchQuery(
    useDeleteProjectMutation.getKey(),
    async () => await useDeleteProjectMutation.fetcher(dataSource, { id })()
  );

  return redirect("/projects");
};

export function ProjectDeletePage() {
  const navigate = useNavigate();

  return (
    <Modal title="Are you sure?" close={() => navigate(-1)}>
      <Form method="post" className="flex flex-col gap-4">
        <input className="btn" type="submit" value="Delete Project" />
      </Form>
    </Modal>
  );
}

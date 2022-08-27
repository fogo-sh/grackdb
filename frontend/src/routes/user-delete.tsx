import { ActionFunction, Form, redirect, useNavigate } from "react-router-dom";
import invariant from "tiny-invariant";
import { Modal } from "~/components/Modal";
import {
  useDeleteUserMutation,
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

  await queryClient.fetchQuery(
    useDeleteUserMutation.getKey(),
    async () => await useDeleteUserMutation.fetcher(dataSource, { id })()
  );

  return redirect("/users");
};

export function UserDeletePage() {
  const navigate = useNavigate();

  return (
    <Modal title="Are you sure?" close={() => navigate(-1)}>
      <Form method="post" className="flex flex-col gap-4">
        <input className="btn" type="submit" value="Delete User" />
      </Form>
    </Modal>
  );
}

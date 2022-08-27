import { ActionFunction, Form, redirect, useNavigate } from "react-router-dom";
import { z } from "zod";
import { Input } from "~/components/Form";
import { Modal } from "~/components/Modal";
import { useCreateUserMutation } from "~/generated/graphql";
import { dataSource, queryClient } from "~/query";

const ActionDataSchema = z.object({
  username: z.string(),
  avatarUrl: z.string(),
});

export const action: ActionFunction = async ({ request }) => {
  const formData = await request.formData();
  const data = ActionDataSchema.parse(Object.fromEntries(formData));

  const resp = await queryClient.fetchQuery(
    useCreateUserMutation.getKey(),
    async () => await useCreateUserMutation.fetcher(dataSource, data)()
  );

  return redirect(`/user/${resp.createUser.username}`);
};

export function UserCreatePage() {
  const navigate = useNavigate();

  return (
    <Modal title="Create User" close={() => navigate("..")}>
      <Form method="post" className="flex flex-col gap-4">
        <Input id="username" display="Username" />
        <Input id="avatarUrl" display="Avatar URL" />
        <input className="btn" type="submit" value="Create User" />
      </Form>
    </Modal>
  );
}

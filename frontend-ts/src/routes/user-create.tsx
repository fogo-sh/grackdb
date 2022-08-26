import { ActionFunction, Form, useNavigate } from "react-router-dom";
import { z } from "zod";
import { Input } from "~/components/Form";
import { Modal } from "~/components/Modal";

const ActionDataSchema = z.object({
  username: z.string(),
  avatarUrl: z.string(),
});

export const action: ActionFunction = async ({ request }) => {
  const formData = await request.formData();
  const data = ActionDataSchema.parse(Object.fromEntries(formData));
  console.log({ data });
};

export function UserCreatePage() {
  const currentUser = undefined; // TODO
  const navigate = useNavigate();

  return (
    <Modal title="Create User" close={() => navigate("..")}>
      <Form method="post" className="flex flex-col gap-4">
        <Input id="username" name="username" display="Username" />
        <Input id="avatarUrl" name="avatarUrl" display="Avatar URL" />
        <input className="btn" type="submit" value="Create User" />
      </Form>
    </Modal>
  );
}

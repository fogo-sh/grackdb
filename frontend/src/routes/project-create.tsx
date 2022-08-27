import { ActionFunction, Form, redirect, useNavigate } from "react-router-dom";
import { z } from "zod";
import { Input } from "~/components/Form";
import { Modal } from "~/components/Modal";
import { useCreateProjectMutation } from "~/generated/graphql";
import { dataSource, queryClient } from "~/query";

const datePreprocess = (arg: unknown) => {
  if (arg === "") return undefined;

  if (typeof arg === "string" || arg instanceof Date) {
    return new Date(arg).toISOString();
  }
};

const DateSchema = z.preprocess(datePreprocess, z.string());
const OptionalDateSchema = z.preprocess(datePreprocess, z.string().optional());

const ActionDataSchema = z.object({
  name: z.string(),
  description: z.string(),
  startDate: DateSchema,
  endDate: OptionalDateSchema,
});

export const action: ActionFunction = async ({ request }) => {
  const formData = await request.formData();
  console.log(Object.fromEntries(formData));
  const data = ActionDataSchema.parse(Object.fromEntries(formData));

  console.log({ data });

  const resp = await queryClient.fetchQuery(
    useCreateProjectMutation.getKey(),
    async () => await useCreateProjectMutation.fetcher(dataSource, data)()
  );

  return redirect(`/project/${resp.createProject.id}`);
};

export function ProjectCreatePage() {
  const navigate = useNavigate();

  // TODO add different project-related things here, like
  // - associate technologies
  // - parent projects
  // - child projects
  // - contribs
  // - sites
  // etc.

  return (
    <Modal title="Create Project" close={() => navigate(-1)}>
      <Form method="post" className="flex flex-col gap-4">
        <Input id="name" display="Name" />
        <Input id="description" display="Description" />
        <Input id="startDate" display="Start Date" type="date" />
        <Input id="endDate" display="End Date" type="date" />
        <input className="btn" type="submit" value="Create Project" />
      </Form>
    </Modal>
  );
}

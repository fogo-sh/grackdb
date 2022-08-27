import { useState } from "react";
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
import { Modal } from "~/components/Modal";
import { Select, SelectItem } from "~/components/Select";
import {
  ProjectTechnologyAssociationType,
  useCreateProjectTechnologyMutation,
  useTechnologiesQuery,
} from "~/generated/graphql";
import { dataSource, queryClient } from "~/query";

const LoaderDataSchema = z.object({
  technologyOptions: z.array(z.object({ id: z.string(), value: z.string() })),
});

type LoaderData = z.infer<typeof LoaderDataSchema>;

export const loader: LoaderFunction = async () => {
  const technologiesQuery = await queryClient.fetchQuery(
    useTechnologiesQuery.getKey(),
    async () => await useTechnologiesQuery.fetcher(dataSource)()
  );

  const technologies =
    technologiesQuery.technologies?.edges?.map((edge) => edge?.node) ?? [];

  const technologyOptions = technologies.map((technology) => ({
    id: technology?.id,
    value: technology?.name,
  }));

  const data: LoaderData = LoaderDataSchema.parse({ technologyOptions });

  return data;
};

const CreateProjectTechnologySchema = z.object({
  type: z.nativeEnum(ProjectTechnologyAssociationType),
  technology: z.number(),
  project: z.number(),
});

export const action: ActionFunction = async ({ request, params }) => {
  const id = params.id;
  invariant(id);

  const thing = params.thing;
  invariant(thing);

  const formData = await request.formData();
  const entries = Object.fromEntries(formData);

  if (thing === "technology") {
    const data = CreateProjectTechnologySchema.parse({
      ...entries,
      project: Number(id),
      technology: Number(entries.technology),
    });
    await queryClient.fetchQuery(
      useCreateProjectTechnologyMutation.getKey(),
      async () =>
        await useCreateProjectTechnologyMutation.fetcher(dataSource, data)()
    );
    return redirect(`/project/${id}`);
  }

  throw new Error("Unknown thing");
};

const associationTypeOptions = Object.entries(
  ProjectTechnologyAssociationType
).map(([k, v]) => ({ id: v, value: k }));

function AssociateTechnology() {
  const { technologyOptions } = useLoaderData() as LoaderData;
  const [selectedAssociationType, setSelectedAssociationType] =
    useState<SelectItem>();
  const [selectedTechnology, setSelectedTechnology] = useState<SelectItem>();

  // TODO create a 'SelectInput' wrapper that makes all of this stuff nicer

  return (
    <>
      <div>
        <p className="mb-1">Association Type</p>
        <Select
          options={associationTypeOptions}
          selected={selectedAssociationType}
          onChange={setSelectedAssociationType}
        />
        <input
          type="hidden"
          name="type"
          value={selectedAssociationType?.id ?? ""}
        />
      </div>
      <div>
        <p className="mb-1">Technology</p>
        <Select
          options={technologyOptions ?? []}
          selected={selectedTechnology}
          onChange={setSelectedTechnology}
        />
        <input
          type="hidden"
          name="technology"
          value={selectedTechnology?.id ?? ""}
        />
      </div>
    </>
  );
}

export function ProjectAssociateThing() {
  const navigate = useNavigate();
  const params = useParams();
  const thing = params.thing;
  const id = params.id;

  return (
    <Modal title="Associate" close={() => navigate(`/project/${id}`)}>
      <Form method="post" className="flex flex-col gap-4">
        {thing === "technology" && <AssociateTechnology />}
        <input className="btn" type="submit" value="Associate" />
      </Form>
    </Modal>
  );
}

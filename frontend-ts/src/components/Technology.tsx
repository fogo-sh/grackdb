import { useEffect } from "react";
import toast from "react-hot-toast";
import { ProjectTechnology, Technology } from "~/types";

import { Circle } from "./Circle";
import { Modal } from "./Modal";
import { Select } from "./Select";

export function TechnologyReference({
  technology,
  children,
}: {
  technology: Technology;
  children: ({
    circle,
    name,
  }: {
    circle: JSX.Element;
    name: string;
  }) => React.ReactNode;
}) {
  const circle = (
    <Circle
      className="h-[0.9rem] ml-2 mr-1"
      fill={technology.colour ?? "gray"}
    />
  );
  const name = technology.name;

  return (
    <div className="flex justify-center items-center">
      {children({ circle, name })}
    </div>
  );
}

export function TechnologiesReference({
  technologies,
}: {
  technologies: { technology: Technology }[];
}) {
  return (
    <div className="flex h-full">
      {technologies.map(({ technology }) => (
        <TechnologyReference key={technology.id} technology={technology}>
          {({ circle, name }) => (
            <>
              {circle} {name}
            </>
          )}
        </TechnologyReference>
      ))}
    </div>
  );
}

const CREATE_PROJECT_TECHNOLOGY_MUTATION = /* GraphQL */ ``;

export function CreateProjectTechnologyModal({
  project,
  refetchProject,
  dialogOpen,
  setDialogOpen,
}) {
  const [
    { data: createProjectTechnologyData, error },
    createProjectTechnology,
  ] = useMutation(CREATE_PROJECT_TECHNOLOGY_MUTATION);
  useErrorNotify(error);

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();

  const onSubmit = (data) => {
    createProjectTechnology({
      ...data,
      project: project.id,
    });
  };

  useEffect(() => {
    if (createProjectTechnologyData) {
      setDialogOpen(false);
      reset();
      toast.success(
        `Associated ${createProjectTechnologyData.createProjectTechnology.type} with ${project.name}`
      );
      refetchProject();
    }
  }, [createProjectTechnologyData]);

  return (
    <Modal
      open={dialogOpen}
      setOpen={setDialogOpen}
      title="Associate Technology"
    >
      <form className="flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)}>
        {/* technology selector */}
        <div>
          <p className="mb-1">Technology</p>
          <Select
            fetching={false}
            options={[]}
            selected={undefined}
            onChange={() => {}}
          />
        </div>
        {/* type selector */}
        <input className="btn" type="submit" value="Associate" />
      </form>
    </Modal>
  );
}

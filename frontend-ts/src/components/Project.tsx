import { format } from "date-fns";
import { Link, useNavigate } from "react-router-dom";
import { useEffect } from "react";
import toast from "react-hot-toast";

import { Modal } from "./Modal";
import { Input } from "./Form";
import { Project } from "~/types";

export function ProjectReference({
  project,
  hasLink = false,
  children,
}: {
  project?: Project | null;
  hasLink: boolean;
  children: ({
    projectName,
  }: {
    projectName: string | JSX.Element;
  }) => React.ReactNode;
}) {
  if (project === undefined || project === null) {
    return <>???</>;
  }

  const projectName = hasLink ? (
    <Link to={`/project/${project.id}`}>{project.name}</Link>
  ) : (
    project.name
  );

  return (
    <div className="flex justify-between my-2 h-6 items-center">
      {children({ projectName })}
    </div>
  );
}

const DATE_FORMAT = "LLLL do, Y";

export function ProjectDates({ project }: { project: Project }) {
  const startDate =
    project.startDate !== null
      ? format(new Date(project.startDate ?? "???"), DATE_FORMAT)
      : "Present";
  const endDate =
    project.endDate !== null
      ? format(new Date(project.endDate ?? "???"), DATE_FORMAT)
      : "Present";

  return (
    <>
      {startDate} - {endDate}
    </>
  );
}

const CREATE_PROJECT_MUTATION = /* GraphQL */ ``;

export function CreateProjectModal({ dialogOpen, setDialogOpen }) {
  const navigate = useNavigate();

  const [{ data: createProjectData, error }, createProject] = useMutation(
    CREATE_PROJECT_MUTATION
  );
  useErrorNotify(error);

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();

  const onSubmit = ({ name, description, startDate, endDate }) => {
    createProject({
      name,
      description: description !== "" ? description : undefined,
      startDate: startDate !== "" ? startDate : undefined,
      endDate: endDate !== "" ? endDate : undefined,
    });
  };

  useEffect(() => {
    if (createProjectData) {
      setDialogOpen(false);
      reset();
      toast.success(`Project ${createProjectData.createProject.name} created!`);
      navigate(`/project/${createProjectData.createProject.id}`);
    }
  });

  return (
    <Modal open={dialogOpen} setOpen={setDialogOpen} title="Create Project">
      <form className="flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)}>
        <Input
          register={register}
          errors={errors}
          id="name"
          name="Name"
          options={{ required: errorMessage.required }}
        />
        <Input
          register={register}
          errors={errors}
          id="description"
          name="Description"
        />
        <Input
          register={register}
          errors={errors}
          id="startDate"
          name="Start Date"
          type="date"
          options={{ required: errorMessage.required, valueAsDate: true }}
        />
        <Input
          register={register}
          errors={errors}
          id="endDate"
          name="End Date"
          type="date"
          options={{ valueAsDate: true }}
        />
        <input className="btn" type="submit" value="Create Project" />
      </form>
    </Modal>
  );
}

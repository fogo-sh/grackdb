import { format } from "date-fns";
import { Link } from "react-router-dom";

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
    <div className="flex justify-between my-2 items-center">
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

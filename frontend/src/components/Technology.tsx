import { Technology } from "~/types";
import { Circle } from "./Circle";

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

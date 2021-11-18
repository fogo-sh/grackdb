import React from "react";

import { Circle } from "./Circle";

export function TechnologyReference({ technology }) {
	return (
		<div className="flex justify-center items-center">
			<Circle className="h-[0.9rem] ml-2 mr-1" fill={technology.colour} />{" "}
			{technology.name}
		</div>
	);
}

export function TechnologiesReference({ technologies }) {
	return (
		<div className="flex h-full">
			{technologies.map(({ technology }) => (
				<TechnologyReference key={technology.id} technology={technology} />
			))}
		</div>
	);
}

import React from "react";

const scalarFilter = ([_key, value]) =>
	value !== null && typeof value !== "object" && !Array.isArray(value);

const nestedScalarFilter = ([_key, value]) =>
	value !== null && typeof value === "object" && !Array.isArray(value);

const subChildrenFilter = ([key, value]) => Array.isArray(value);

const DEPTH_INC = 0;

function View({ data }) {
	const renderChildren = (children, depth) => {
		const { __typename, id, ...rest } = children;

		if (!id) {
			throw new Error(`missing 'id' field on ${__typename}'`);
		}

		const key = `${__typename}:${id}`;

		const scalars = Object.entries(rest).filter(scalarFilter);
		const nestedScalars = Object.entries(rest).filter(nestedScalarFilter);
		const subChildren = Object.entries(rest).filter(subChildrenFilter);

		return (
			<div key={key} className="pt-1">
				{scalars.map(([scalarKey, scalarKeyValue]) => (
					<p key={`${key}:${scalarKey}`}>
						<span className="italic font-light">{scalarKey}</span>:{" "}
						{scalarKeyValue}
					</p>
				))}
				{nestedScalars.map(([nestedKey, { id, __typename, ...rest }]) => (
					<div key={`${key}:${nestedKey}`}>
						<p className="italic font-light">{nestedKey}:</p>
						{Object.entries(rest).map(
							([nestedNestedKey, nestedNestedValue]) => (
								<div
									key={`${key}:${nestedKey}:${nestedNestedKey}`}
									className="ml-3"
								>
									<p>
										<span className="italic font-light">{nestedNestedKey}</span>
										: {nestedNestedValue}
									</p>
								</div>
							)
						)}
					</div>
				))}
				{subChildren.map(([subChildKey, subChildKeyValue]) => (
					<div key={`${key}:${subChildKey}`} className="ml-3">
						<p className="italic font-light">{subChildKey}:</p>
						<div className="ml-3">
							{subChildKeyValue.map((subChild) =>
								renderChildren(subChild, depth + DEPTH_INC)
							)}
						</div>
					</div>
				))}
			</div>
		);
	};

	if (!data) {
		return <div>Loading...</div>;
	}

	return Object.entries(data).map(([key, value]) => {
		return (
			<div key={key} className="">
				<p className="text-xl">{key}</p>
				{value.edges.map(({ node }) => renderChildren(node, 1))}
			</div>
		);
	});
}

export default View;

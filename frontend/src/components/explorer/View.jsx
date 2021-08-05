import React from "react";

function View({ data }) {
	const renderChildren = (children, depth) => {
		const { __typename, id, ...rest } = children;

		if (!id) {
			throw new Error(`missing 'id' field on ${__typename}'`);
		}

		const scalars = Object.entries(rest).filter(
			([_key, value]) =>
				value !== null && typeof value !== "object" && !Array.isArray(value)
		);

		const nestedScalars = Object.entries(rest).filter(
			([_key, value]) =>
				value !== null && typeof value === "object" && !Array.isArray(value)
		);

		const subChildren = Object.entries(rest).filter(([key, value]) =>
			Array.isArray(value)
		);

		const key = `${__typename}:${id}`;

		return (
			<div key={key}>
				{scalars.map(([scalarKey, scalarKeyValue]) => (
					<p key={`${key}:${scalarKey}`}>
						{scalarKey}: {scalarKeyValue}
					</p>
				))}
				{nestedScalars.map(([nestedKey, { id, __typename, ...rest }]) => (
					<div key={`${key}:${nestedKey}`}>
						<p>{nestedKey}:</p>
						{Object.entries(rest).map(
							([nestedNestedKey, nestedNestedValue]) => (
								<div
									key={`${key}:${nestedKey}:${nestedNestedKey}`}
									style={{ marginLeft: `${depth}rem` }}
								>
									<p>
										{nestedNestedKey}: {nestedNestedValue}
									</p>
								</div>
							)
						)}
					</div>
				))}
				{subChildren.map(([subChildKey, subChildKeyValue]) => (
					<div key={`${key}:${subChildKey}`}>
						<p>{subChildKey}:</p>
						<div style={{ marginLeft: `${depth + 1}rem` }}>
							{subChildKeyValue.map((subChild) =>
								renderChildren(subChild, depth + 1)
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
				<p className="text-xl mb-2">{key}</p>
				{value.edges.map(({ node }) => renderChildren(node, 1))}
			</div>
		);
	});
}

export default View;

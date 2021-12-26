import React from "react";

export function Input({ register, id, name, options }) {
	return (
		<div className="flex flex-col">
			<label htmlFor={id} className="mb-1">
				{name}
			</label>
			<input
				className="inp"
				id={id}
				autoComplete="off"
				{...register(id, options)}
			/>
		</div>
	);
}

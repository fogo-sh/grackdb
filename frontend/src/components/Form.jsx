import React from "react";
import { ErrorMessage } from "@hookform/error-message";

export function Input({ register, errors, id, name, options, type = "text" }) {
	const error = errors[id];

	return (
		<div className="flex flex-col">
			<label htmlFor={id} className="mb-1">
				{name}
			</label>
			<input
				className="inp"
				id={id}
				autoComplete="off"
				type={type}
				{...register(id, options)}
			/>
			<ErrorMessage
				errors={errors}
				name={id}
				render={({ message }) => (
					<p className="text-sm italic text-red-600 mt-0.5">{message}</p>
				)}
			/>
		</div>
	);
}

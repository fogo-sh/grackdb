import React from "react";

export function DeleteButton({ onClick }) {
	return (
		<button className="btn" onClick={onClick}>
			<span className="mx-1">x</span>
		</button>
	);
}

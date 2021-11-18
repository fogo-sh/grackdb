import React from "react";

export function Circle({ className = "", fill = "black" }) {
	return (
		<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg" className={className}>
			<circle cx="50" cy="50" r="50" fill={fill} />
		</svg>
	);
}

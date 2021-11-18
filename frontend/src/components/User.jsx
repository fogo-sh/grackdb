import React from "react";
import { Link } from "react-router-dom";

export function UserReference({ user, hasLink = false }) {
	return (
		<div className="my-2">
			{hasLink ? (
				<Link to={`/user/${user.username}`}>{user.username}</Link>
			) : (
				user.username
			)}
		</div>
	);
}

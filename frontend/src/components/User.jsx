import React from "react";
import { Link } from "react-router-dom";

export function UserReference({ user, hasLink = false, children }) {
	const userName = hasLink ? (
		<Link to={`/user/${user.username}`}>{user.username}</Link>
	) : (
		user.username
	);

	return <div className="my-2">{children({ userName })}</div>;
}

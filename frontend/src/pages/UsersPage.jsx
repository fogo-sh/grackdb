import React from "react";
import { useQuery } from "urql";

import { UserReference } from "../components/User";

const USERS_QUERY = `
query Users {
	users {
		edges {
			node {
				id
				username
				avatarUrl
			}
		}
	}
}
`;

export function UsersPage() {
	const [{ fetching, data }] = useQuery({
		query: USERS_QUERY,
	});

	if (fetching) {
		return null;
	}

	const users = data.users.edges.map((edge) => edge.node);

	return (
		<>
			<h2>Users</h2>
			<div className="mx-2">
				{users.map((user) => (
					<UserReference key={user.id} user={user} hasLink />
				))}
			</div>
		</>
	);
}

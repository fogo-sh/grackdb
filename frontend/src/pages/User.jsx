import React from "react";
import { useQuery } from "urql";
import { useParams } from "react-router-dom";

import { UserReference } from "../components/User";

const USERS_BY_USERNAME_QUERY = `
query UsersByUsername($username: String!) {
	users(where: {username: $username}) {
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

export function UserPage() {
	const params = useParams();

	const [{ fetching, data }] = useQuery({
		query: USERS_BY_USERNAME_QUERY,
		variables: { username: params.username },
	});

	if (fetching) {
		return null;
	}

	const user = data.users.edges[0].node;

	return <UserReference user={user} />;
}

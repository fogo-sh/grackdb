import React, { useState } from "react";
import { useQuery } from "urql";

import { CreateUserModal, UserReference } from "../components/User";
import { useAuth } from "../providers/AuthProvider";

const USERS_QUERY = /* GraphQL */ `
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
	const [dialogOpen, setDialogOpen] = useState(false);

	const { currentUser } = useAuth();

	const [{ fetching, data }] = useQuery({
		query: USERS_QUERY,
	});

	if (fetching) {
		return null;
	}

	const users = data.users.edges.map((edge) => edge.node);

	return (
		<>
			<div className="flex justify-between items-center">
				<h2 className="h-full my-0">Users</h2>
				{currentUser && (
					<>
						<button className="btn h-1/2" onClick={() => setDialogOpen(true)}>
							Create
						</button>
						<CreateUserModal
							dialogOpen={dialogOpen}
							setDialogOpen={setDialogOpen}
						/>
					</>
				)}
			</div>
			<div className="mx-2">
				{users.map((user) => (
					<UserReference key={user.id} user={user} hasLink>
						{({ userName }) => userName}
					</UserReference>
				))}
			</div>
		</>
	);
}

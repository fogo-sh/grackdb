import React, { createContext, useContext } from "react";
import { useQuery } from "urql";

const AuthContext = createContext();

export const useAuth = () => useContext(AuthContext);

const CURRENT_USER_QUERY = /* GraphQL */ `
	query CurrentUser {
		currentUser {
			id
			username
			avatarUrl
		}
	}
`;

export function AuthProvider({ children }) {
	const [{ data, refetch }, reexecuteQuery] = useQuery({
		query: CURRENT_USER_QUERY,
	});

	const refreshCurrentUser = () => {
		reexecuteQuery({ requestPolicy: "network-only" });
	};

	return (
		<AuthContext.Provider
			value={{
				currentUser: data?.currentUser ?? undefined,
				refreshCurrentUser,
			}}
		>
			{children}
		</AuthContext.Provider>
	);
}

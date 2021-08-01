import React from "react";
import { useQuery } from "urql";

const UsersQuery = `
{
  users {
    edges {
      node {
        username
        avatarUrl
      }
    }
  }
}
`;

function Users() {
  const [{ data, fetching, error }, reexecuteQuery] = useQuery({
    query: UsersQuery,
  });

  if (fetching) return <p>Loading...</p>;
  if (error) return <p>Oh no... {error.message}</p>;

  return (
    <ul>
      {data.users.edges.map(({ node }) => (
        <li key={node.username}>{node.username}</li>
      ))}
    </ul>
  );
}

export default Users;
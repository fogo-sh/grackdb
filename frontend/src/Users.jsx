import React from "react";
import { useQuery } from "urql";

const UsersQuery = `
{
  users {
    edges {
      node {
        id
        username
        avatarUrl
        discordAccounts {
          id
          username
          discriminator
        }
        githubAccounts {
          id
          username
          organizationMemberships {
            id
            role
            organization {
              name
              displayName
            }
          }
        }
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
    <div className="p-2">
      <p className="text-xl pb-5">Users</p>
      <ul>
        {data.users.edges.map(({ node }) => (
          <li key={node.id} className="pb-4">
            <p className="text-lg">{node.username}</p>
            {node.discordAccounts.length > 0 && (
              <ul className="mt-2 pl-3">
                {node.discordAccounts.map(({ id, username, discriminator }) => (
                  <li key={id}>
                    <i class="fab fa-discord" /> {username}#{discriminator}
                  </li>
                ))}
              </ul>
            )}
            {node.githubAccounts.length > 0 && (
              <ul className="mt-2 pl-3">
                {node.githubAccounts.map(
                  ({ id, username, organizationMemberships }) => (
                    <li key={id}>
                      <i class="fab fa-github-square" /> {username}
                      {organizationMemberships.length > 0 && (
                        <ul className="mt-2 pl-3">
                          Organizations:
                          {organizationMemberships.map(
                            ({ id, role, organization }) => (
                              <li key={id}>
                                {organization.name} ({role})
                              </li>
                            )
                          )}
                        </ul>
                      )}
                    </li>
                  )
                )}
              </ul>
            )}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Users;

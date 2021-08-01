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
            account {
              username
            }
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
      <p>Users</p>
      <ul class="list-disc">
        {data.users.edges.map(({ node }) => (
          <li key={node.id} className="pb-4">
            {node.username}
            {node.discordAccounts.length > 0 && (
              <ul className="mt-2 pl-3">
                Discord:
                {node.discordAccounts.map(({ id, username, discriminator }) => (
                  <li key={id}>
                    {username}#{discriminator}
                  </li>
                ))}
              </ul>
            )}
            {node.githubAccounts.length > 0 && (
              <ul className="mt-2 pl-3">
                GitHub:
                {node.githubAccounts.map(
                  ({ id, username, organizationMemberships }) => (
                    <li key={id}>
                      {username}
                      {organizationMemberships.length > 0 && (
                        <ul className="mt-2 pl-3">
                          Organizations:
                          {organizationMemberships.map(
                            ({ id, role, account, organization }) => (
                              <li key={id}>
                                {role}
                                {account.username}
                                {organization.name}
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

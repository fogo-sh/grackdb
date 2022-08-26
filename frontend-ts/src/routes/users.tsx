import { Link, LoaderFunction, Outlet, useLoaderData } from "react-router-dom";
import { z } from "zod";
import { UserSchema } from "~/types";
import { dataSource, queryClient } from "~/query";
import { useUsersQuery } from "~/generated/graphql";
import { UserReference } from "~/components/User";

const LoaderDataSchema = z.object({
  users: z.array(UserSchema),
});

type LoaderData = z.infer<typeof LoaderDataSchema>;

export const loader: LoaderFunction = async () => {
  const usersQuery = await queryClient.fetchQuery(
    useUsersQuery.getKey(),
    async () => await useUsersQuery.fetcher(dataSource)()
  );

  const users = usersQuery.users?.edges?.map((edge) => edge?.node);

  const data: LoaderData = LoaderDataSchema.parse({ users });
  return data;
};

export function UsersPage() {
  const { users } = useLoaderData() as LoaderData;
  const currentUser = true; // TODO

  return (
    <>
      <div className="flex justify-between items-center">
        <h2 className="h-full my-0">Users</h2>
        {currentUser && (
          <Link to="./create">
            <button className="btn h-1/2">Create</button>
          </Link>
        )}
      </div>
      <div className="mx-2">
        {users.map((user) => (
          <UserReference key={user.id} user={user} hasLink>
            {({ userName }) => userName}
          </UserReference>
        ))}
      </div>
      <Outlet />
    </>
  );
}

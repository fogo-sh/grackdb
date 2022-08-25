import { useLoaderData } from "react-router-dom";
import { z } from "zod";
import { UserSchema } from "~/types";
import { dataSource, queryClient } from "~/query";
import { useUsersQuery } from "~/generated/graphql";
import { UserReference } from "~/components/User";

const LoaderDataSchema = z.object({
  users: z.array(UserSchema),
});

type LoaderData = z.infer<typeof LoaderDataSchema>;

export async function loader() {
  const homepage = await queryClient.fetchQuery(
    useUsersQuery.getKey(),
    async () => await useUsersQuery.fetcher(dataSource)()
  );

  const users = homepage.users?.edges?.map((edge) => edge?.node);

  const data: LoaderData = LoaderDataSchema.parse({ users });
  return data;
}

export function Users() {
  const { users } = useLoaderData() as LoaderData;
  const currentUser = undefined;

  return (
    <>
      <div className="flex justify-between items-center">
        <h2 className="h-full my-0">Users</h2>
        {currentUser && (
          <>
            <button className="btn h-1/2" onClick={() => alert(true)}>
              Create
            </button>
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

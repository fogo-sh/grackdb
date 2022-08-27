import { useState } from "react";
import { FaCrown } from "react-icons/fa";
import {
  ActionFunction,
  Form,
  LoaderFunction,
  redirect,
  useLoaderData,
  useNavigate,
} from "react-router-dom";
import { z } from "zod";
import { UserSchema } from "~/types";
import { useAssumableUsersQuery } from "~/generated/graphql";
import { dataSource, queryClient } from "~/query";
import { Select, SelectItem } from "~/components/Select";
import { Modal } from "~/components/Modal";

/*
  const navigate = useNavigate();
  const { refreshCurrentUser } = useAuth();

  const [{ fetching, data }] = useQuery({
    query: DEVELOPMENT_ASSUME_USER_QUERY,
  });

  const [{ data: assumeDevelopmentUserData, error }, assumeDevelopmentUser] =
    useMutation(DEVELOPMENT_ASSUME_USER_MUTATION);
  useErrorNotify(error);

  useEffect(() => {
    if (data === undefined) {
      return;
    }

    const users = data.users.edges.map(({ node: { id, username } }) => ({
      id,
      value: username,
    }));

    setOptions(users);
    setSelectedUser(users[0]);
  }, [data]);

  const assume = useCallback(() => {
    assumeDevelopmentUser({ id: selectedUser.id });
  }, [selectedUser]);

  useEffect(() => {
    if (assumeDevelopmentUserData === undefined) {
      return;
    }

    refreshCurrentUser();
    navigate("/");
  }, [assumeDevelopmentUserData]);

  if (!selectedUser) {
    return null;
  }
*/

const LoaderDataSchema = z.object({
  options: z.array(z.object({ id: z.string(), value: z.string() })),
});

type LoaderData = z.infer<typeof LoaderDataSchema>;

export const loader: LoaderFunction = async () => {
  const assumeUsersQuery = await queryClient.fetchQuery(
    useAssumableUsersQuery.getKey(),
    async () => await useAssumableUsersQuery.fetcher(dataSource)()
  );

  const users = assumeUsersQuery.users?.edges?.map((edge) => edge?.node) ?? [];

  const options = users.map((user) => ({
    id: user?.id,
    value: user?.username,
  }));

  const data: LoaderData = LoaderDataSchema.parse({ options });

  return data;
};

const ActionDataSchema = z.object({});

export const action: ActionFunction = async ({ request }) => {
  const formData = await request.formData();
  console.log(Object.fromEntries(formData));
  // const data = ActionDataSchema.parse(Object.fromEntries(formData));

  return redirect("/");
};

export function AssumeUserPage() {
  const { options } = useLoaderData() as LoaderData;

  const navigate = useNavigate();

  const [selectedUser, setSelectedUser] = useState<SelectItem>();

  return (
    <Modal close={() => navigate("..")}>
      <p className="text-center mb-3 italic font-semibold">
        ⭐ Super Secret Developer Assume User Tool ⭐
      </p>

      <Form
        method="post"
        className="flex justify-center items-center gap-2 h-10"
      >
        <Select
          options={options ?? []}
          selected={selectedUser}
          onChange={setSelectedUser}
        />
        <button
          className="btn flex w-[6rem] my-3 h-full items-center"
          type="submit"
        >
          <div className="flex items-center mx-auto">
            <FaCrown className="mr-1" /> Assume
          </div>
        </button>
      </Form>
    </Modal>
  );
}

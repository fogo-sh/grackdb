import { createContext, useContext } from "react";
import { LoaderFunction, Outlet, useLoaderData } from "react-router-dom";
import { z } from "zod";
import { Header } from "~/components/Header";
import { useCurrentUserQuery } from "~/generated/graphql";
import { dataSource, queryClient } from "~/query";
import { User, UserSchema } from "~/types";

const LoaderDataSchema = z.object({ currentUser: UserSchema.nullable() });

type LoaderData = z.infer<typeof LoaderDataSchema>;

export const loader: LoaderFunction = async () => {
  const { currentUser } = await queryClient.fetchQuery(
    useCurrentUserQuery.getKey(),
    async () => await useCurrentUserQuery.fetcher(dataSource)()
  );

  const data: LoaderData = LoaderDataSchema.parse({ currentUser });
  return data;
};

export const CurrentUserContext = createContext<User | null>(null);

export const useCurrentUser = () => useContext(CurrentUserContext);

export function RootPage() {
  const { currentUser } = useLoaderData() as LoaderData;

  return (
    <div className="w-11/12 sm:w-2/3 lg:w-7/12 xl:w-1/2 2xl:w-5/12 mx-auto my-3">
      <Header currentUser={currentUser} />
      <hr className="my-3" />
      <CurrentUserContext.Provider value={currentUser}>
        <Outlet />
      </CurrentUserContext.Provider>
    </div>
  );
}

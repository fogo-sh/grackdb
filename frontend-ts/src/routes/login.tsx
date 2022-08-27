import { FaDiscord, FaGithub } from "react-icons/fa";
import { Link, LoaderFunction, Outlet, useLoaderData } from "react-router-dom";
import { z } from "zod";
import { AuthProviderType, useLoginQuery } from "~/generated/graphql";
import { dataSource, queryClient } from "~/query";
import { AuthProviderSchema } from "~/types";

const LoaderDataSchema = z.object({
  availableAuthProviders: z.array(AuthProviderSchema),
  developmentMode: z.boolean(),
});

type LoaderData = z.infer<typeof LoaderDataSchema>;

export const loader: LoaderFunction = async () => {
  const { availableAuthProviders, developmentMode } =
    await queryClient.fetchQuery(
      useLoginQuery.getKey(),
      async () => await useLoginQuery.fetcher(dataSource)()
    );

  const data: LoaderData = LoaderDataSchema.parse({
    availableAuthProviders,
    developmentMode,
  });

  return data;
};

export function LoginPage() {
  const { developmentMode, availableAuthProviders } =
    useLoaderData() as LoaderData;

  return (
    <>
      <h1 className="text-center">Login</h1>

      {availableAuthProviders.length === 0 && (
        <p className="text-center italic">no available auth providers!</p>
      )}

      <div className="flex flex-col items-center">
        {availableAuthProviders.map((provider) => (
          <a key={provider.type} href={provider.url}>
            <button className="btn flex w-[7rem] my-2">
              <div className="flex items-center mx-auto">
                {provider.type === AuthProviderType.Discord && (
                  <>
                    <FaDiscord className="mr-1" /> Discord
                  </>
                )}
                {provider.type === AuthProviderType.Github && (
                  <>
                    <FaGithub className="mr-1" /> GitHub{" "}
                  </>
                )}
              </div>
            </button>
          </a>
        ))}
        <Link to="./assume">
          <button className="btn flex w-[10rem] mt-4 my-2">
            <div className="flex items-center mx-auto">⭐ Assume User</div>
          </button>
        </Link>
      </div>
      <Outlet />
    </>
  );
}

import React, { useCallback, useEffect, useState } from "react";
import { useQuery, useMutation } from "urql";
import { FaDiscord, FaGithub, FaCrown } from "react-icons/fa";
import { useAuth } from "../providers/AuthProvider";
import { useNavigate } from "react-router-dom";
import { Select } from "../components/Select";

function DevelopmentAssumeUser() {
  const [options, setOptions] = useState(null);
  const [selectedUser, setSelectedUser] = useState(null);

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

  return (
    <>
      <p className="text-center mb-3 italic font-semibold">
        ⭐ Super Secret Developer Assume User Tool ⭐
      </p>

      <div className="flex justify-center items-center gap-2 h-10">
        <Select
          fetching={fetching}
          options={options}
          selected={selectedUser}
          onChange={setSelectedUser}
        />
        <button
          className="btn flex w-[6rem] my-3 h-full items-center"
          onClick={assume}
        >
          <div className="flex items-center mx-auto">
            <FaCrown className="mr-1" /> Assume
          </div>
        </button>
      </div>
    </>
  );
}

export function LoginPage() {
  const [{ fetching, data }] = useQuery({
    query: LOGIN_QUERY,
  });

  if (fetching) {
    return null;
  }

  const developmentMode = data.developmentMode;

  const availableAuthProviders = Object.fromEntries(
    data.availableAuthProviders.map(({ type, ...authProvider }) => [
      type.toLowerCase(),
      authProvider,
    ])
  );

  return (
    <>
      <h1 className="text-center">Login</h1>

      {data.availableAuthProviders.length === 0 && (
        <p className="text-center italic">no available auth providers!</p>
      )}

      <div className="flex flex-col items-center">
        {availableAuthProviders.discord && (
          <a href={availableAuthProviders.discord.url}>
            <button className="btn flex w-[6rem] my-2">
              <div className="flex items-center mx-auto">
                <FaDiscord className="mr-1" /> Discord
              </div>
            </button>
          </a>
        )}

        {availableAuthProviders.github && (
          <a href={availableAuthProviders.github.url}>
            <button className="btn flex w-[6rem] my-2">
              <div className="flex items-center mx-auto">
                <FaGithub className="mr-1" /> GitHub
              </div>
            </button>
          </a>
        )}
      </div>

      {developmentMode && (
        <>
          <br className="my-6" />
          <DevelopmentAssumeUser />
        </>
      )}
    </>
  );
}

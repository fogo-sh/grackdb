import { Link } from "react-router-dom";

import React from "react";
import { User } from "~/types";

export function UserReference({
  user,
  hasLink = false,
  children,
}: {
  user: User;
  hasLink: boolean;
  children: ({
    userName,
  }: {
    userName: string | JSX.Element;
  }) => React.ReactNode;
}) {
  const userName = hasLink ? (
    <Link to={`/user/${user.username}`}>{user.username}</Link>
  ) : (
    user.username
  );

  return <div className="my-2">{children({ userName })}</div>;
}

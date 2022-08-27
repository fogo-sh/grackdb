import { Link } from "react-router-dom";
import { User } from "~/types";

import grackUrl from "~/images/grack.png";

export function Header({ currentUser }: { currentUser: User | null }) {
  return (
    <>
      <div className="flex items-center justify-between">
        <div className="w-[3.5rem] flex">
          <img src={grackUrl} className="h-16 mb-1 mx-auto" />
        </div>
        {currentUser && <span>Logged in as {currentUser.username}</span>}
      </div>
      <div className="flex flex-col sm:flex-row gap-1 sm:gap-0 justify-between">
        <div className="flex flex-col sm:flex-row gap-1 sm:gap-3">
          <Link to="/">GrackDB</Link>
          <Link to="/users">Users</Link>
          <Link to="/projects">Projects</Link>
        </div>
        <div className="flex flex-col sm:flex-row gap-1 sm:gap-3">
          <a href="/playground/">GraphQL Playground</a>
          {currentUser ? (
            <Link to="/logout">Logout</Link>
          ) : (
            <Link to="/login">Login</Link>
          )}
        </div>
      </div>
    </>
  );
}

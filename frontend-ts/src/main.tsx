import React from "react";
import ReactDOM from "react-dom/client";
import { DataBrowserRouter, Route } from "react-router-dom";
import "./index.css";

import ErrorPage from "./error-page";
import { RootPage, loader as rootLoader } from "./routes/root";
import { IndexPage, loader as indexLoader } from "./routes/index";
import { UsersPage, loader as usersLoader } from "./routes/users";
import {
  UserCreatePage,
  action as userCreateAction,
} from "./routes/user-create";
import { UserPage, loader as userLoader } from "./routes/user";
import {
  UserDeletePage,
  action as userDeleteAction,
} from "./routes/user-delete";
import {
  UserDeleteThing,
  action as userDeleteThingAction,
} from "./routes/user-delete-thing";
import { ProjectsPage, loader as projectsLoader } from "./routes/projects";
import { ProjectPage, loader as projectLoader } from "./routes/project";
import {
  ProjectDeletePage,
  action as projectDeleteAction,
} from "./routes/project-delete";
import { LoginPage, loader as loginLoader } from "./routes/login";
import {
  AssumeUserPage,
  loader as assumeUserLoader,
  action as assumeUserAction,
} from "./routes/login-assume-user";
import { LogoutPage, loader as logoutLoader } from "./routes/logout";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <DataBrowserRouter>
      <Route
        path="/"
        element={<RootPage />}
        loader={rootLoader}
        errorElement={<ErrorPage />}
      >
        <Route path="" element={<IndexPage />} loader={indexLoader} />
        <Route path="users/" element={<UsersPage />} loader={usersLoader}>
          <Route
            path="create"
            element={<UserCreatePage />}
            action={userCreateAction}
          />
        </Route>
        <Route path="user/:username" element={<UserPage />} loader={userLoader}>
          <Route
            path="delete"
            element={<UserDeletePage />}
            action={userDeleteAction}
          />
          <Route
            path="delete/:thing/:thingId"
            element={<UserDeleteThing />}
            action={userDeleteThingAction}
          />
        </Route>
        <Route
          path="projects"
          element={<ProjectsPage />}
          loader={projectsLoader}
        />
        <Route
          path="project/:id"
          element={<ProjectPage />}
          loader={projectLoader}
        >
          <Route
            path="delete"
            element={<ProjectDeletePage />}
            action={projectDeleteAction}
          />
        </Route>
        <Route path="login" element={<LoginPage />} loader={loginLoader}>
          <Route
            path="assume"
            element={<AssumeUserPage />}
            loader={assumeUserLoader}
            action={assumeUserAction}
          />
        </Route>
        <Route path="logout" element={<LogoutPage />} loader={logoutLoader} />
      </Route>
    </DataBrowserRouter>
  </React.StrictMode>
);

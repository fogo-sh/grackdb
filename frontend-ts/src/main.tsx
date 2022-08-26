import React from "react";
import ReactDOM from "react-dom/client";
import { DataBrowserRouter, Route } from "react-router-dom";
import "./index.css";

import ErrorPage from "./error-page";
import { RootPage } from "./routes/root";
import { IndexPage, loader as indexLoader } from "./routes/index";
import { UsersPage, loader as usersLoader } from "./routes/users";
import {
  UserCreatePage,
  action as userCreateAction,
} from "./routes/user-create";
import { UserPage, loader as userLoader } from "./routes/user";
import { ProjectsPage, loader as projectsLoader } from "./routes/projects";
import { ProjectPage, loader as projectLoader } from "./routes/project";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <DataBrowserRouter>
      <Route path="/" element={<RootPage />} errorElement={<ErrorPage />}>
        <Route path="" element={<IndexPage />} loader={indexLoader} />
        <Route path="users/" element={<UsersPage />} loader={usersLoader}>
          <Route
            path="create"
            element={<UserCreatePage />}
            action={userCreateAction}
          />
        </Route>
        <Route
          path="user/:username"
          element={<UserPage />}
          loader={userLoader}
        />
        <Route
          path="projects"
          element={<ProjectsPage />}
          loader={projectsLoader}
        />
        <Route
          path="project/:id"
          element={<ProjectPage />}
          loader={projectLoader}
        />
      </Route>
    </DataBrowserRouter>
  </React.StrictMode>
);

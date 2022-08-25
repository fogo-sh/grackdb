import React from "react";
import ReactDOM from "react-dom/client";
import { DataBrowserRouter, Route } from "react-router-dom";
import "./index.css";

import Root from "./routes/root";
import ErrorPage from "./error-page";
import Index, { loader as indexLoader } from "./routes/index";
import Users, { loader as usersLoader } from "./routes/users";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <DataBrowserRouter>
      <Route path="/" element={<Root />} errorElement={<ErrorPage />}>
        <Route path="" element={<Index />} loader={indexLoader} />
        <Route path="users" element={<Users />} loader={usersLoader} />
      </Route>
    </DataBrowserRouter>
  </React.StrictMode>
);

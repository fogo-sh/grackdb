import React from "react";
import { Provider } from "urql";
import { BrowserRouter, Routes, Route, Link } from "react-router-dom";

import { client } from "./graphql";
import { HomepagePage } from "./pages/HomepagePage";
import { ProjectsPage } from "./pages/ProjectsPage";
import { ProjectPage } from "./pages/ProjectPage";
import { UsersPage } from "./pages/UsersPage";
import { UserPage } from "./pages/UserPage";

import grack from "./grack.png";

function App() {
	return (
		<BrowserRouter>
			<Provider value={client}>
				<div className="w-11/12 sm:w-2/3 lg:w-7/12 xl:w-1/2 2xl:w-5/12 mx-auto my-3">
					<div className="w-[3.5rem]">
						<img src={grack} className="h-16 mb-1 mx-auto" />
					</div>
					<div className="flex justify-between">
						<div className="flex gap-3">
							<Link to="/">GrackDB</Link>
							<Link to="/users">Users</Link>
							<Link to="/projects">Projects</Link>
						</div>
						<div className="flex gap-3">
							<a href="/playground">GraphQL Playground</a>
						</div>
					</div>
					<hr className="my-3" />
					<Routes>
						<Route path="/" element={<HomepagePage />} />
						<Route path="/projects" element={<ProjectsPage />} />
						<Route path="/project/:projectId" element={<ProjectPage />} />
						<Route path="/users" element={<UsersPage />} />
						<Route path="/user/:username" element={<UserPage />} />
					</Routes>
				</div>
			</Provider>
		</BrowserRouter>
	);
}

export default App;

import React from "react";
import { Provider } from "urql";
import { BrowserRouter, Routes, Route, Link } from "react-router-dom";

import { client } from "./graphql";
import { HomepagePage } from "./pages/HomepagePage";
import { ProjectsPage } from "./pages/ProjectsPage";
import { ProjectPage } from "./pages/ProjectPage";
import { UsersPage } from "./pages/UsersPage";
import { UserPage } from "./pages/UserPage";
import { LoginPage } from "./pages/LoginPage";
import { LogoutPage } from "./pages/LogoutPage";
import Cookies from "js-cookie";

import grack from "./grack.png";
import { AuthProvider, useAuth } from "./providers/AuthProvider";

function Header() {
	const { currentUser } = useAuth();

	return (
		<>
			<div className="flex items-center justify-between">
				<div className="w-[3.5rem] flex">
					<img src={grack} className="h-16 mb-1 mx-auto" />
				</div>
				{currentUser && <span>Logged in as {currentUser.username}</span>}
			</div>
			<div className="flex justify-between">
				<div className="flex gap-3">
					<Link to="/">GrackDB</Link>
					<Link to="/users">Users</Link>
					<Link to="/projects">Projects</Link>
				</div>
				<div className="flex gap-3">
					<a href="/playground">GraphQL Playground</a>
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

function AppRoutes() {
	return (
		<Routes>
			<Route path="/" element={<HomepagePage />} />
			<Route path="/projects" element={<ProjectsPage />} />
			<Route path="/project/:projectId" element={<ProjectPage />} />
			<Route path="/users" element={<UsersPage />} />
			<Route path="/user/:username" element={<UserPage />} />
			<Route path="/login" element={<LoginPage />} />
			<Route path="/logout" element={<LogoutPage />} />
		</Routes>
	);
}

function App() {
	return (
		<BrowserRouter>
			<Provider value={client}>
				<AuthProvider>
					<div className="w-11/12 sm:w-2/3 lg:w-7/12 xl:w-1/2 2xl:w-5/12 mx-auto my-3">
						<Header />
						<hr className="my-3" />
						<AppRoutes />
					</div>
				</AuthProvider>
			</Provider>
		</BrowserRouter>
	);
}

export default App;

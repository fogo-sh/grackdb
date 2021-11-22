import React, { useCallback, useEffect, useState } from "react";
import { useQuery, useMutation } from "urql";
import { Listbox } from "@headlessui/react";
import {
	FaDiscord,
	FaGithub,
	FaCrown,
	FaCaretDown,
	FaCheck,
} from "react-icons/fa";
import { useAuth } from "../providers/AuthProvider";
import { useNavigate } from "react-router-dom";

const LOGIN_QUERY = `
query Login {
  availableAuthProviders {
    url
    type
  }
	developmentMode
	users {
		edges {
			node {
				id
				username
			}
		}
	}
}
`;

const DEVELOPMENT_ASSUME_USER_QUERY = `
query AssumableUsers {
	users {
		edges {
			node {
				id
				username
			}
		}
	}
}
`;

const DEVELOPMENT_ASSUME_USER_MUTATION = `
mutation AssumeDevelopmentUser($id: ID!) {
  assumeDevelopmentUser(id: $id) {
    id
  }
}
`;

function DevelopmentAssumeUser() {
	const [options, setOptions] = useState(null);
	const [selectedUser, setSelectedUser] = useState(null);

	const navigate = useNavigate();
	const { refreshCurrentUser } = useAuth();

	const [{ fetching, data }] = useQuery({
		query: DEVELOPMENT_ASSUME_USER_QUERY,
	});

	const [{ data: assumeDevelopmentUserData }, assumeDevelopmentUser] =
		useMutation(DEVELOPMENT_ASSUME_USER_MUTATION);

	useEffect(() => {
		if (data === undefined) {
			return;
		}

		const users = data.users.edges.map((edge) => edge.node);
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

	if (fetching || !selectedUser) {
		return null;
	}

	return (
		<>
			<p className="text-center mb-3 italic font-semibold">
				⭐ Super Secret Developer Assume User Tool ⭐
			</p>

			<div className="flex justify-center items-center gap-2 h-10">
				<Listbox value={selectedUser} onChange={setSelectedUser}>
					<div className="relative min-w-[15rem] h-full">
						<Listbox.Button className="relative w-full h-full border border-gray-500 rounded-sm pl-3 pr-2 py-1 flex items-center justify-between mx-auto">
							{selectedUser.username}{" "}
							<FaCaretDown className="ml-1 opacity-50" />
						</Listbox.Button>
						<Listbox.Options className="absolute w-full mt-1 overflow-auto text-base bg-white rounded-m max-h-60 border border-gray-500">
							{options.map((option) => (
								<Listbox.Option key={option.id} value={option}>
									{({ selected, active }) => (
										<div
											className={`${
												active ? "bg-gray-100" : ""
											} flex items-center px-2.5 py-1.5`}
										>
											<span
												className={`${
													selected ? "font-medium" : "font-normal"
												} block truncate`}
											>
												{option.username}
											</span>
											{selected ? (
												<FaCheck
													className="ml-2 opacity-50"
													aria-hidden="true"
												/>
											) : null}
										</div>
									)}
								</Listbox.Option>
							))}
						</Listbox.Options>
					</div>
				</Listbox>

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

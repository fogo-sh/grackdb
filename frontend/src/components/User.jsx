import React from "react";
import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";

import { errorMessage } from "../consts";
import { Modal } from "./Modal";
import { Input } from "./Form";
import { useMutation } from "urql";
import { useEffect } from "react";
import toast from "react-hot-toast";
import { useErrorNotify } from "../hooks/useErrorNotify";

export function UserReference({ user, hasLink = false, children }) {
	const userName = hasLink ? (
		<Link to={`/user/${user.username}`}>{user.username}</Link>
	) : (
		user.username
	);

	return <div className="my-2">{children({ userName })}</div>;
}

const CREATE_USER_MUTATION = /* GraphQL */ `
	mutation CreateUser($username: String!, $avatarUrl: String) {
		createUser(input: { username: $username, avatarUrl: $avatarUrl }) {
			id
			username
			avatarUrl
		}
	}
`;

export function CreateUserModal({ dialogOpen, setDialogOpen }) {
	const navigate = useNavigate();

	const [{ data: createUserData, error }, createUser] =
		useMutation(CREATE_USER_MUTATION);
	useErrorNotify(error);

	const {
		register,
		handleSubmit,
		reset,
		formState: { errors },
	} = useForm();

	const onSubmit = ({ username, avatarUrl }) => {
		createUser({
			username,
			avatarUrl: avatarUrl !== "" ? avatarUrl : undefined,
		});
	};

	useEffect(() => {
		if (createUserData) {
			setDialogOpen(false);
			reset();
			toast.success(`User ${createUserData.createUser.username} created!`);
			navigate(`/user/${createUserData.createUser.username}`);
		}
	}, [createUserData]);

	return (
		<Modal open={dialogOpen} setOpen={setDialogOpen} title="Create User">
			<form className="flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)}>
				<Input
					register={register}
					errors={errors}
					id="username"
					name="Username"
					options={{ required: errorMessage.required }}
				/>
				<Input
					register={register}
					errors={errors}
					id="avatarUrl"
					name="Avatar URL"
				/>
				<input className="btn" type="submit" value="Create User" />
			</form>
		</Modal>
	);
}

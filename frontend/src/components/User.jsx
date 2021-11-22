import React from "react";
import { useForm } from "react-hook-form";
import { Link } from "react-router-dom";

import { Modal } from "./Modal";
import { Input } from "./Form";
import { useMutation } from "urql";
import { useEffect } from "react";
import toast from "react-hot-toast";

export function UserReference({ user, hasLink = false, children }) {
	const userName = hasLink ? (
		<Link to={`/user/${user.username}`}>{user.username}</Link>
	) : (
		user.username
	);

	return <div className="my-2">{children({ userName })}</div>;
}

const CREATE_USER_MUTATION = `
mutation CreateUser($username: String!) {
  createUser(input: {username: $username}) {
    id
    username
  }
}
`;

export function CreateUserModal({ dialogOpen, setDialogOpen }) {
	const [{ data: createUserData }, createUser] =
		useMutation(CREATE_USER_MUTATION);

	const {
		register,
		handleSubmit,
		reset,
		formState: { errors },
	} = useForm();

	const onSubmit = (data) => createUser(data);

	useEffect(() => {
		if (createUserData) {
			setDialogOpen(false);
			reset();
			toast.success(`User ${createUserData.createUser.username} created!`);
		}
	}, [createUserData]);

	return (
		<Modal open={dialogOpen} setOpen={setDialogOpen} title="Create User">
			<form className="flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)}>
				<Input
					register={register}
					id="username"
					name="Username"
					options={{ required: true }}
				/>
				{errors.exampleRequired && <span>This field is required</span>}
				<input className="btn" type="submit" value="Create User" />
			</form>
		</Modal>
	);
}

import React, { useEffect } from "react";
import { useForm } from "react-hook-form";
import toast from "react-hot-toast";
import { FaDiscord } from "react-icons/fa";
import { useMutation } from "urql";
import { errorMessage } from "../consts";
import { useErrorNotify } from "../hooks/useErrorNotify";

import { Input } from "./Form";
import { Modal } from "./Modal";

export function DiscordAccountReference({ discordAccount }) {
	return (
		<div className="flex my-2 items-center">
			<FaDiscord className="mr-1" /> {discordAccount.username}
			<span className="opacity-40">#{discordAccount.discriminator}</span>
		</div>
	);
}

const CREATE_DISCORD_ACCOUNT_MUTATION = `
mutation CreateDiscordAccount(
  $discordId: String!
  $username: String!
  $discriminator: String!
  $owner: Int!
) {
  createDiscordAccount(
    input: {
      discordId: $discordId
      username: $username
      discriminator: $discriminator,
      owner: $owner
    }
  ) {
    id
    username
  }
}
`;

export function CreateDiscordAccountModal({
	user,
	refetchUser,
	dialogOpen,
	setDialogOpen,
}) {
	const [{ data: createDiscordAccountData, error }, createDiscordAccount] =
		useMutation(CREATE_DISCORD_ACCOUNT_MUTATION);
	useErrorNotify(error);

	const {
		register,
		handleSubmit,
		reset,
		formState: { errors },
	} = useForm();

	const onSubmit = (data) => createDiscordAccount({ ...data, owner: user.id });

	useEffect(() => {
		if (createDiscordAccountData) {
			setDialogOpen(false);
			reset();
			toast.success(
				`Associated ${createDiscordAccountData.createDiscordAccount.username} with ${user.username}`
			);
			refetchUser();
		}
	}, [createDiscordAccountData]);

	return (
		<Modal
			open={dialogOpen}
			setOpen={setDialogOpen}
			title="Associate Discord Account"
		>
			<form className="flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)}>
				<Input
					register={register}
					errors={errors}
					id="username"
					name="Discord Username"
					options={{ required: errorMessage.required }}
				/>
				<Input
					register={register}
					errors={errors}
					id="discriminator"
					name="Discord Discriminator"
					options={{ required: errorMessage.required }}
				/>
				<Input
					register={register}
					errors={errors}
					id="discordId"
					name="Discord ID"
					options={{ required: errorMessage.required }}
				/>
				<input
					className="btn"
					type="submit"
					value="Associate Discord Account"
				/>
			</form>
		</Modal>
	);
}

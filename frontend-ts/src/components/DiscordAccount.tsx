import { useEffect } from "react";
import toast from "react-hot-toast";
import { FaDiscord } from "react-icons/fa";
import { DiscordAccount } from "~/types";

import { Input } from "./Form";
import { Modal } from "./Modal";

export function DiscordAccountReference({
  discordAccount,
}: {
  discordAccount: DiscordAccount;
}) {
  return (
    <div className="flex my-2 items-center">
      <FaDiscord className="mr-1" /> {discordAccount.username}
      <span className="opacity-40">#{discordAccount.discriminator}</span>
    </div>
  );
}

const CREATE_DISCORD_ACCOUNT_MUTATION = /* GraphQL */ ``;

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
        <input
          className="btn"
          type="submit"
          value="Associate Discord Account"
        />
      </form>
    </Modal>
  );
}

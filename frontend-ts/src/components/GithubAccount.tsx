import { useEffect } from "react";
import toast from "react-hot-toast";
import { FaExternalLinkAlt, FaGithub } from "react-icons/fa";
import { Input } from "./Form";
import { Modal } from "./Modal";

export function GithubAccountReference({ githubAccount, hasLink = false }) {
  return (
    <div className="flex my-2 items-center">
      <FaGithub className="mr-1" />
      {hasLink ? (
        <>
          <a
            href={`https://github.com/${githubAccount.username}`}
            className="flex items-center"
          >
            {githubAccount.username}
            <FaExternalLinkAlt className="ml-1 text-xs" />
          </a>
        </>
      ) : (
        githubAccount.username
      )}
    </div>
  );
}

const CREATE_GITHUB_ACCOUNT_MUTATION = /* GraphQL */ ``;

export function CreateGithubAccountModal({
  user,
  refetchUser,
  dialogOpen,
  setDialogOpen,
}) {
  const [{ data: createGithubAccountData, error }, createGithubAccount] =
    useMutation(CREATE_GITHUB_ACCOUNT_MUTATION);
  useErrorNotify(error);

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();

  const onSubmit = (data) => createGithubAccount({ ...data, owner: user.id });

  useEffect(() => {
    if (createGithubAccountData) {
      setDialogOpen(false);
      reset();
      toast.success(
        `Associated ${createGithubAccountData.createGithubAccount.username} with ${user.username}`
      );
      refetchUser();
    }
  }, [createGithubAccountData]);

  return (
    <Modal
      open={dialogOpen}
      setOpen={setDialogOpen}
      title="Associate GitHub Account"
    >
      <form className="flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)}>
        <Input
          register={register}
          errors={errors}
          id="username"
          name="GitHub Username"
          options={{ required: errorMessage.required }}
        />
        <input className="btn" type="submit" value="Associate GitHub Account" />
      </form>
    </Modal>
  );
}

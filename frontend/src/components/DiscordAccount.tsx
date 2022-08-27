import { FaDiscord } from "react-icons/fa";
import { DiscordAccount } from "~/types";

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

import React from "react";
import { FaDiscord } from "react-icons/fa";

export function DiscordAccountReference({ discordAccount }) {
	return (
		<div className="flex my-2 items-center">
			<FaDiscord className="mr-1" /> {discordAccount.username}
			<span className="opacity-40">#{discordAccount.discriminator}</span>
		</div>
	);
}

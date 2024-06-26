import { apiClient } from "~/utils/api-client";
import type { Channel } from "./channel";

type GetChannelsResponse = {
	channels: Channel[];
};

export const getChannels = async () => {
	const { channels } = (await apiClient
		.get("channels")
		.json()) as GetChannelsResponse;
	return channels;
};

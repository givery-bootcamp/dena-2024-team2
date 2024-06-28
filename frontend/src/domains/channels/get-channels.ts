import { apiClient } from "~/utils/api-client";
import type { Channel } from "./channel";

type GetChannelsResponse = {
	channels: Channel[];
};

export const getChannels = (serverId?: number) => async () => {
	if (!serverId) {
		return [];
	}
	const { channels } = (await apiClient("timeline")
		.get(`servers/${serverId}/channels`)
		.json()) as GetChannelsResponse;
	return channels;
};

import { apiClient } from "~/utils/api-client";

type PostChannelRequest = {
	name: string;
};

type PostChannelResponse = {
	id: number;
	serverId: number;
	name: string;
	createdAt: string;
};

export const postChannel =
	(serverId: number) => async (data: PostChannelRequest) => {
		const channel = (await apiClient("timeline")
			.post(`servers/${serverId}/channels`, {
				json: data,
			})
			.json()) as PostChannelResponse;
		return channel;
	};

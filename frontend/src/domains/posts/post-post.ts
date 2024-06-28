import { apiClient } from "~/utils/api-client";

type PostPostRequest = {
	content: string;
};

export const postPost =
	(serverId: number, channelId: number) => async (data: PostPostRequest) => {
		const post = await apiClient("timeline")
			.post(`servers/${serverId}/channels/${channelId}/posts`, {
				json: data,
			})
			.json();
		return post;
	};

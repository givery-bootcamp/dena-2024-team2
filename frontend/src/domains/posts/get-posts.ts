import { apiClient } from "~/utils/api-client";
import type { Post } from "./post";

type GetPostsResponse = {
	posts: Post[];
};

export const getPosts = (channelId?: number) => async () => {
	if (!channelId) {
		return [];
	}
	const { posts } = (await apiClient("timeline")
		.get(`channels/${channelId}/posts`)
		.json()) as GetPostsResponse;
	return posts;
};

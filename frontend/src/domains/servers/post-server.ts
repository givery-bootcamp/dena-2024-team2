import { apiClient } from "~/utils/api-client";

type PostPostRequest = {
	name: string;
};

export const postServer = async (data: PostPostRequest) => {
	const post = await apiClient("timeline")
		.post("servers", {
			json: data,
		})
		.json();
	return post;
};

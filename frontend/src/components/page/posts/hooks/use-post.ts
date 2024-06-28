import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import { useState } from "react";
import { postPost } from "~/domains/posts";

export const usePost = (serverId: number, channelId: number) => {
	const [content, setContent] = useState("");
	const navigate = useNavigate();

	const { mutate, status } = useMutation({
		mutationFn: postPost(serverId, channelId),
		onSuccess() {
			location.reload();
		},
	});
};

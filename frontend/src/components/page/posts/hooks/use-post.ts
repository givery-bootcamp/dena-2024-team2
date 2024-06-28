import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import { postPost } from "~/domains/posts";

export const usePost = (serverId: number, channelId: number) => {
	const [content, setContent] = useState("");

	const { mutate, status } = useMutation({
		mutationFn: postPost(serverId, channelId),
		onSuccess() {
			location.reload();
		},
	});

	const handleChangeContent = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
		setContent(e.target.value);
	};

	const handleClickSubmit = () => {
		mutate({ content });
	};

	return {
		status,
		handleChangeContent,
		handleClickSubmit,
	};
};

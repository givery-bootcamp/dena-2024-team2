import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import { postChannel } from "~/domains/channels";

export const useChannel = (serverId: number) => {
	const [name, setName] = useState("");

	const { mutate, status } = useMutation({
		mutationFn: postChannel(serverId),
		onSuccess() {
			location.reload();
		},
	});

	const handleChangeName = (e: React.ChangeEvent<HTMLInputElement>) => {
		setName(e.target.value);
	};

	const handleCreateChannel = () => {
		mutate({ name });
	};

	return {
		status,
		handleChangeName,
		handleCreateChannel,
	};
};

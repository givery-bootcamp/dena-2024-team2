import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import { postServer } from "~/domains/servers";

export const useServer = () => {
	const [name, setName] = useState("");
	const [icon, setIcon] = useState("");

	const { mutate, status } = useMutation({
		mutationFn: postServer,
		onSuccess() {
			location.reload();
		},
	});

	const handleChangeName = (e: React.ChangeEvent<HTMLInputElement>) => {
		setName(e.target.value);
	};
	const handleChangeIcon = (e: React.ChangeEvent<HTMLInputElement>) => {
		setIcon(e.target.value);
	};

	const handleCreateSever = () => {
		mutate({ name, icon });
	};

	return {
		status,
		handleChangeName,
		handleChangeIcon,
		handleCreateSever,
	};
};

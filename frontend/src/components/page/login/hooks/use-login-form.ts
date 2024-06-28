import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import { useState } from "react";
import { postSignin } from "~/domains/auth";

export const useLoginForm = () => {
	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const navigate = useNavigate();

	const { mutate, status } = useMutation({
		mutationFn: postSignin,
		onSuccess(data) {
			localStorage.setItem("user", JSON.stringify(data));
			navigate({ to: "/servers" });
		},
	});

	const handleChangeUsername = (e: React.ChangeEvent<HTMLInputElement>) => {
		setUsername(e.target.value);
	};

	const handleChangePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
		setPassword(e.target.value);
	};

	const handleClickLogin = () => {
		mutate({ name: username, password });
	};

	return {
		status,
		handleChangeUsername,
		handleChangePassword,
		handleClickLogin,
	};
};

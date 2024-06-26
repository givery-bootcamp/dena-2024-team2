import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import { postSignin } from "~/domains/auth";

export const useLoginForm = () => {
	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");

	const { mutate, status } = useMutation({
		mutationFn: postSignin,
		onSuccess(data) {
			// FIXME: cookieに入るようになったら削除
			localStorage.setItem("TMP_TOKEN", data.accessToken);
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

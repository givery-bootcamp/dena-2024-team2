import { apiClient } from "~/utils/api-client";

type SigninRequest = {
	name: string;
	password: string;
};

// FIXME: ここの型は今後変更になる
type SigninResponse = {
	accessToken: string;
};

export const postSignin = async (data: SigninRequest) => {
	const token = (await apiClient
		.post("signin", {
			json: data,
		})
		.json()) as SigninResponse;
	return token;
};

import { apiClient } from "~/utils/api-client";

type SigninRequest = {
	name: string;
	password: string;
};

// FIXME: ここの型は今後変更になる
type SigninResponse = {
	id: number;
	name: string;
};

export const postSignin = async (data: SigninRequest) => {
	const token = (await apiClient("user")
		.post("user/signin", {
			json: data,
		})
		.json()) as SigninResponse;
	return token;
};

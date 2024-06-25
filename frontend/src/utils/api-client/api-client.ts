import ky from "ky";

export const apiClient = ky
	.create({ prefixUrl: process.env.API_BASE_URL })
	.extend({
		hooks: {
			beforeRequest: [
				(request) => {
					request.headers.set(
						"Authorization",
						// FIXME: 一時的に localStorage から取得するようにしている
						// 本来はset-cookieするので、そうなったら修正する
						`Bearer ${localStorage.getItem("TMP_TOKEN")}`,
					);
				},
			],
		},
	});

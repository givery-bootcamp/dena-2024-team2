import ky from "ky";
import { convertKeyCaseRecursive } from "../convert-key-recursive/convert-key-case-recursive";

export const apiClient = ky
	.create({ prefixUrl: import.meta.env.VITE_API_BASE_URL })
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
			afterResponse: [
				// camelCaseをsnake_caseに変換する処理
				async (_request, _options, response) => {
					if (
						response.headers.get("content-type")?.includes("application/json")
					) {
						const jsonRes = await response.json();
						const modifiedResponse = convertKeyCaseRecursive(jsonRes, "camel");
						return new Response(JSON.stringify(modifiedResponse));
					}
					return response;
				},
			],
		},
	});

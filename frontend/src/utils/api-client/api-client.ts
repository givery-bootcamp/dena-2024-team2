import ky from "ky";
import { convertKeyCaseRecursive } from "../convert-key-recursive/convert-key-case-recursive";

const prefixUrlMap = {
	timeline: import.meta.env.VITE_TIMELINE_API_BASE_URL,
	user: import.meta.env.VITE_USER_API_BASE_URL,
};

export const apiClient = (service: "timeline" | "user") =>
	ky
		.create({ prefixUrl: prefixUrlMap[service], credentials: "include" })
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
						if (response.status === 401) {
							// 401の場合はログイン画面にリダイレクト
							// 本来はrouteのloaderでfetchして、401の場合はログイン画面にリダイレクトするようにするのが正しそう…？
							location.href = "/login";
						}

						if (
							response.headers.get("content-type")?.includes("application/json")
						) {
							const jsonRes = await response.json();
							const modifiedResponse = convertKeyCaseRecursive(
								jsonRes,
								"camel",
							);
							return new Response(JSON.stringify(modifiedResponse));
						}
						return response;
					},
				],
			},
		});

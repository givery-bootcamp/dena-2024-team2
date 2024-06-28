import { apiClient } from "~/utils/api-client";
import type { Server } from "./server";

type GetServersResponse = {
	servers: Server[];
};

export const getServers = async () => {
	const { servers } = (await apiClient("timeline")
		.get("servers")
		.json()) as GetServersResponse;
	return servers;
};

import type { Server } from "~/domains/servers";

export const mock = {
	servers: [
		{
			id: "1",
			ownerId: "1",
			name: "My Server",
			icon: "https://via.placeholder.com/50",
		},
		{
			id: "2",
			ownerId: "1",
			name: "Another Server",
			icon: "https://via.placeholder.com/50",
		},
	],
} as { servers: Server[] };

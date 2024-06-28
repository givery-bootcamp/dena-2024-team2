import type { Channel } from "~/domains/channels";
import type { Server } from "~/domains/servers";

export const mock = {
	channels: [
		{
			id: 1,
			serverId: 1,
			name: "general",
			createdAt: "2021-01-01T00:00:00Z",
			updatedAt: "2021-01-01T00:00:00Z",
		},
		{
			id: 2,
			serverId: 1,
			name: "random",
			createdAt: "2021-01-01T00:00:00Z",
			updatedAt: "2021-01-01T00:00:00Z",
		},
	],
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
} as { channels: Channel[]; servers: Server[] };

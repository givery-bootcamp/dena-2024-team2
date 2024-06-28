import type { Channel } from "~/domains/channels";
import type { Post } from "~/domains/posts";
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
	posts: [
		{
			id: 1,
			channelId: 1,
			user: {
				id: 1,
				name: "John Doe",
			},
			content: "Hello, world!",
			createdAt: "2021-01-01T00:00:00Z",
		},
		{
			id: 2,
			channelId: 2,
			user: {
				id: 2,
				name: "Jane Smith",
			},
			content: "This is a test post.",
			createdAt: "2021-01-01T00:00:00Z",
		},
		{
			id: 3,
			channelId: 1,
			user: {
				id: 1,
				name: "John Doe",
			},
			content: "This is another test post.",
			createdAt: "2021-01-01T00:00:00Z",
		},
		{
			id: 4,
			channelId: 2,
			user: {
				id: 2,
				name: "Jane Smith",
			},
			content: "This is yet another test post.",
			createdAt: "2021-01-01T00:00:00Z",
		},
		{
			id: 5,
			channelId: 1,
			user: {
				id: 1,
				name: "John Doe",
			},
			content: "This is the last test post.",
			createdAt: "2021-01-01T00:00:00Z",
		},
		{
			id: 6,
			channelId: 2,
			user: {
				id: 2,
				name: "Jane Smith",
			},
			content: "This is the final test post.",
			createdAt: "2021-01-01T00:00:00Z",
		},
		{
			id: 7,
			channelId: 1,
			user: {
				id: 1,
				name: "John Doe",
			},
			content: "This is the very last test post.",
			createdAt: "2021-01-01T00:00:00Z",
		},
		{
			id: 8,
			channelId: 2,
			user: {
				id: 2,
				name: "Jane Smith",
			},
			content: "This is the very final test post.",
			createdAt: "2021-01-01T00:00:00Z",
		},
		{
			id: 9,
			channelId: 1,
			user: {
				id: 1,
				name: "John Doe",
			},
			content: "This is the very very last test post.",
			createdAt: "2021-01-01T00:00:00Z",
		},
		{
			id: 10,
			channelId: 2,
			user: {
				id: 2,
				name: "Jane Smith",
			},
			content: "This is the very very final test post.",
			createdAt: "2021-01-01T00:00:00Z",
		},
	],
} as { channels: Channel[]; servers: Server[]; posts: Post[] };

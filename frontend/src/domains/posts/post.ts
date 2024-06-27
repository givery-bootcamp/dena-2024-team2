import type { User } from "../users";

export type Post = {
	id: number;
	channelId: number;
	user: User;
	content: string;
	createdAt: Date;
};

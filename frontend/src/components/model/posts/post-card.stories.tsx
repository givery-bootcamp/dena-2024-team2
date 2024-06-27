import type { Meta, StoryFn } from "@storybook/react";

import type { ComponentProps } from "react";
import { PostCard } from "./post-card";

export default {
	title: "Components/Model/Posts/PostCard",
	component: PostCard,
	tags: ["autodocs"],
} as Meta;

const Template: StoryFn<ComponentProps<typeof PostCard>> = (args) => (
	<PostCard {...args} />
);

export const ListOnly = Template.bind({});
ListOnly.args = {
	post: {
		id: 1,
		channelId: 1,
		user: {
			id: 1,
			name: "John Doe",
		},
		content: "This is a post",
		createdAt: "2021-01-01T00:00:00.000Z",
	},
};

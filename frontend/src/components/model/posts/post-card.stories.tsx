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

export const Default = Template.bind({});
Default.args = {
	post: {
		id: 1,
		channelId: 1,
		userId: 1,
		userName: "John Doe",
		content: "This is a post",
		createdAt: "2024-06-26T17:31:15+09:00",
	},
};

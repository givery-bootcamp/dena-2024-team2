import type { Meta, StoryFn } from "@storybook/react";

import type { ComponentProps } from "react";
import { PostList } from "./post-list";

export default {
	title: "Components/Model/Posts/PostList",
	component: PostList,
	tags: ["autodocs"],
} as Meta;

const Template: StoryFn<ComponentProps<typeof PostList>> = (args) => (
	<PostList {...args} />
);

export const Default = Template.bind({});
Default.args = {
	posts: [
		{
			id: 1,
			channelId: 1,
			user: {
				id: 1,
				name: "John Doe",
			},
			content: "This is a post",
			createdAt: "2024-06-26T17:31:15+09:00",
		},
		{
			id: 2,
			channelId: 1,
			user: {
				id: 2,
				name: "Jane Doe",
			},
			content: "This is another post",
			createdAt: "2024-06-26T17:31:15+09:00",
		},
	],
};

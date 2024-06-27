import type { Meta, StoryFn } from "@storybook/react";

import { fn } from "@storybook/test";
import type { ComponentProps } from "react";
import { ServerList } from "./server-list";

export default {
	title: "Components/Model/Server/ServerList",
	component: ServerList,
	tags: ["autodocs"],
} as Meta;

const Template: StoryFn<ComponentProps<typeof ServerList>> = (args) => (
	<ServerList {...args} />
);

export const ListOnly = Template.bind({});
ListOnly.args = {
	servers: [
		{
			id: "1",
			ownerId: "owner1",
			name: "Server 1",
			icon: "https://via.placeholder.com/50",
		},
		{
			id: "2",
			ownerId: "owner2",
			name: "Server 2",
			icon: "https://via.placeholder.com/50",
		},
	],
};

export const WithPlusButton = Template.bind({});
WithPlusButton.args = {
	...ListOnly.args,
	onClickPlus: fn(),
};

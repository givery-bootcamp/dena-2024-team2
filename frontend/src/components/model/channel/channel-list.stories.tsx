import type { Meta, StoryFn } from "@storybook/react";

import { fn } from "@storybook/test";
import type { ComponentProps } from "react";
import { ChannelList } from "./channel-list";

export default {
	title: "Components/Model/Channel/ChannelList",
	component: ChannelList,
	tags: ["autodocs"],
} as Meta;

const Template: StoryFn<ComponentProps<typeof ChannelList>> = (args) => (
	<ChannelList {...args} />
);

export const ListOnly = Template.bind({});
ListOnly.args = {
	channels: [
		{
			id: 1,
			serverId: 1,
			name: "channel 1",
			createdAt: "2021-01-01T00:00:00.000Z",
			updatedAt: "2021-01-01T00:00:00.000Z",
		},
		{
			id: 2,
			serverId: 2,
			name: "channel 2",
			createdAt: "2021-01-01T00:00:00.000Z",
			updatedAt: "2021-01-01T00:00:00.000Z",
		},
	],
};

export const WithPlusButton = Template.bind({});
WithPlusButton.args = {
	...ListOnly.args,
	onClickPlus: fn(),
};

export const CurrentChannel = Template.bind({});
CurrentChannel.args = {
	...ListOnly.args,
	currentChannelId: 1,
};

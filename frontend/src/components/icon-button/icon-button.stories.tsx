import type { Meta, StoryFn } from "@storybook/react";

import type { ComponentProps } from "react";
import { IconButton } from "./icon-button";

export default {
	title: "Components/IconButton",
	component: IconButton,
} as Meta;

const Template: StoryFn<ComponentProps<typeof IconButton>> = (args) => (
	<IconButton {...args} />
);

export const Default = Template.bind({});
Default.args = {
	icon: "plus",
	size: "md",
	color: "primary",
};

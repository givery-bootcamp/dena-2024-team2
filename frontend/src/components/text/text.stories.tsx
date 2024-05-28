import type { Meta, StoryFn } from "@storybook/react";

import type { ComponentProps } from "react";
import { Text } from "./text";

export default {
	title: "Components/Text",
	component: Text,
	tags: ["autodocs"],
} as Meta;

const Template: StoryFn<ComponentProps<typeof Text>> = (args) => (
	<Text {...args} />
);

export const Default = Template.bind({});
Default.args = {
	as: "p",
	size: "md",
	children: "Sample Text",
};

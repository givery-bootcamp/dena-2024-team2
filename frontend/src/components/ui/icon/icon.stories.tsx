import type { Meta, StoryFn } from "@storybook/react";

import type { ComponentProps } from "react";
import { Icon } from "./icon";

export default {
	title: "Components/Icon",
	component: Icon,
	tags: ["autodocs"],
	argTypes: {
		icon: {
			control: "text",
		},
		size: {
			control: {
				type: "radio",
			},
			options: ["sm", "md", "lg"],
		},
		color: {
			control: {
				type: "radio",
			},
			options: ["primary", "secondary"],
		},
	},
} as Meta;

const Template: StoryFn<ComponentProps<typeof Icon>> = (args) => (
	<Icon {...args} />
);

export const Default = Template.bind({});
Default.args = {
	icon: "plus",
	size: "md",
	color: "primary",
};

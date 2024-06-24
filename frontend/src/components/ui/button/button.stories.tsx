import type { Meta, StoryFn } from "@storybook/react";

import type { ComponentProps } from "react";
import { Button } from "./button";

export default {
	title: "Components/Button",
	component: Button,
	tags: ["autodocs"],
} as Meta;

const Template: StoryFn<ComponentProps<typeof Button>> = (args) => (
	<Button {...args} />
);

export const Default = Template.bind({});
Default.args = {
	size: "md",
	variant: "fill",
	color: "primary",
	children: "Button",
};

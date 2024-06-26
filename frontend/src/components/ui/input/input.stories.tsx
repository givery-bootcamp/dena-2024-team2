import type { Meta, StoryFn } from "@storybook/react";

import type { ComponentProps } from "react";
import { Input } from "./input";

export default {
	title: "Components/Input",
	component: Input,
	tags: ["autodocs"],
} as Meta;

const Template: StoryFn<ComponentProps<typeof Input>> = (args) => (
	<Input {...args} />
);

export const Default = Template.bind({});
Default.args = {
	placeholder: "Enter your name",
};

import type { Meta, StoryFn } from "@storybook/react";

import type { ComponentProps } from "react";
import { Textarea } from "./textarea";

export default {
	title: "Components/Textarea",
	component: Textarea,
	tags: ["autodocs"],
} as Meta;

const Template: StoryFn<ComponentProps<typeof Textarea>> = (args) => (
	<Textarea {...args} />
);

export const Default = Template.bind({});
Default.args = {
	rows: 3,
	placeholder: "Type here...",
};

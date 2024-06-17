import type { Meta, StoryFn } from "@storybook/react";

import { Header } from "./header";

export default {
	title: "Components/Header",
	component: Header,
	tags: ["autodocs"],
} as Meta;

const Template: StoryFn = () => <Header />;

export const Default = Template.bind({});
Default.args = {
	// Add any necessary props here
};

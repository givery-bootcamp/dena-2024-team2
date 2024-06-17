import type { Meta, StoryFn } from "@storybook/react";

import type { ComponentProps } from "react";
import { IconImage } from "./icon-image";

export default {
	title: "Components/IconImage",
	component: IconImage,
	argTypes: {
		src: {
			control: "text",
			description: "画像のURL",
		},
	},
} as Meta;

const Template: StoryFn<ComponentProps<typeof IconImage>> = (args) => (
	<IconImage {...args} />
);

export const Default = Template.bind({});
Default.args = {
	src: "https://via.placeholder.com/150",
};

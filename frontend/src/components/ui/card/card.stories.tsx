import type { Meta, StoryFn } from "@storybook/react";

import type { ComponentProps } from "react";
import { Card } from "./card";

export default {
	title: "Components/Card",
	component: Card,
	tags: ["autodocs"],
} as Meta;

const Template: StoryFn<ComponentProps<typeof Card>> = (args) => (
	<Card header={args.header} body={args.body} />
);

export const Default = Template.bind({});
Default.args = {
	header: <h3>Header</h3>,
	body: (
		<div>
			<p>this is body</p>
			<p>this is body</p>
			<p>this is body</p>
			<p>this is body</p>
		</div>
	),
};

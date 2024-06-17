export default {
	stories: ["../src/**/*.mdx", "../**/*.stories.@(js|jsx|mjs|ts|tsx)"],
	addons: ["@storybook/addon-links", "@storybook/addon-essentials"],

	core: {
		builder: "@storybook/builder-vite",
	},

	framework: {
		name: "@storybook/react-vite",
		options: {},
	},

	async viteFinal(config) {
		// Merge custom configuration into the default config
		const { mergeConfig } = await import("vite");

		return mergeConfig(config, {
			// Add dependencies to pre-optimization
			optimizeDeps: {
				include: ["storybook-dark-mode"],
			},
		});
	},

	docs: {
		autodocs: true,
	},
};

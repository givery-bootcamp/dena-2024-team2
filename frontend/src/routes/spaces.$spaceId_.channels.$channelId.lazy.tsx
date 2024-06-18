import { createLazyFileRoute } from "@tanstack/react-router";

export const Route = createLazyFileRoute(
	"/spaces/$spaceId/channels/$channelId",
)({
	component: () => <div>Hello /spaces/$spaceId/channels/$channelId!</div>,
});

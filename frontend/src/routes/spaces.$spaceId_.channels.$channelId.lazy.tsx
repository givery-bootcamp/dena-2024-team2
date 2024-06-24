import { createLazyFileRoute } from "@tanstack/react-router";
import { ChannelDetailPage } from "~/components/page";

export const Route = createLazyFileRoute(
	"/spaces/$spaceId/channels/$channelId",
)({
	component: () => <ChannelDetailPage />,
});

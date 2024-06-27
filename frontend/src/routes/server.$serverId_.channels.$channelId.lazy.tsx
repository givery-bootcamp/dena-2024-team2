import { createLazyFileRoute } from "@tanstack/react-router";
import { ChannelDetailPage } from "~/components/page";

export const Route = createLazyFileRoute(
	"/server/$serverId/channels/$channelId",
)({
	component: () => <ChannelDetailPage />,
});

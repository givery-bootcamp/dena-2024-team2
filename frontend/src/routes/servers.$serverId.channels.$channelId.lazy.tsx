import { createLazyFileRoute } from "@tanstack/react-router";
import { ChannelDetailPage } from "~/components/page";

export const Route = createLazyFileRoute(
	"/servers/$serverId/channels/$channelId",
)({
	component: () => <ChannelDetailPage />,
});

import { createLazyFileRoute } from "@tanstack/react-router";
import { Posts } from "~/components/page/posts";

export const Route = createLazyFileRoute(
	"/servers/$serverId/channels/$channelId",
)({
	component: () => <Posts />,
});
